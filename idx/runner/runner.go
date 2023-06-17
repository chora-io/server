package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
)

// Runner runs continuous background processes.
type Runner struct {
	ctx               context.Context
	backoffDuration   time.Duration
	backoffMaxRetries uint64
	client            client.Client
	waitGroup         sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, cfg config.Config, client client.Client) Runner {
	return Runner{
		ctx:               ctx,
		backoffDuration:   cfg.BackoffDuration,
		backoffMaxRetries: cfg.BackoffMaxRetries,
		client:            client,
	}
}

// ProcessFunc is the function used to advance the process.
type ProcessFunc func(ctx context.Context, c client.Client, chainId string) error

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(chainId string, processName string, processFunc ProcessFunc) {
	// add process to wait group
	r.waitGroup.Add(1)

	go func() {
		// decrement wait group counter on exit
		defer r.waitGroup.Done()

		defer fmt.Println("stopping process", processName)

		// set and initialize backoff
		backoffDuration := r.backoffDuration
		backoffMaxRetries := r.backoffMaxRetries
		backoffRetryCount := uint64(0)

		fmt.Println("starting process", processName)

		for {
			// set process start time
			processStart := time.Now()

			// execute process function
			err := processFunc(r.ctx, r.client, chainId)

			// set process duration
			processDuration := time.Since(processStart)

			fmt.Println("process duration", processName, processDuration.String())

			if err != nil {
				fmt.Println("process error", processName, err.Error())

				// update retry count
				backoffRetryCount++
			}

			// exit on max retries
			if backoffRetryCount == backoffMaxRetries {
				fmt.Println("maximum retries", processName, backoffMaxRetries)
				return // exit process
			} else if backoffRetryCount > 0 {
				fmt.Println("retry count", processName, backoffRetryCount)
			}

			// wait for the backoff duration to continue or context done to exit
			select {
			case <-time.After(backoffDuration):
				fmt.Println("backing off", processName, backoffDuration.String())
				continue // continue process
			case <-r.ctx.Done():
				return // exit process
			}
		}
	}()
}

// Close closes the runner.
func (r *Runner) Close() {
	fmt.Println("finishing processes")

	// wait for processes to finish
	r.waitGroup.Wait()

	fmt.Println("shutting down client")

	// close client
	err := r.client.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("shutdown complete")
}
