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
	// Context is the context used for cancellation and shutdown of the runner.
	ctx context.Context

	// backoffDuration is the duration between advancing processes.
	backoffDuration time.Duration

	// backoffMaxRetries
	backoffMaxRetries uint64

	// client is the client used to interact with the database and network.
	client client.Client

	// waitGroup tracks running processes and is used to coordinate a shutdown.
	waitGroup sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, cfg config.Config, client client.Client) Runner {
	if ctx == nil {
		ctx = context.Background()
	}
	return Runner{
		ctx:               ctx,
		backoffDuration:   cfg.IdxBackoffDuration,
		backoffMaxRetries: cfg.IdxBackoffMaxRetries,
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

		// log stopping process on exit
		defer func() {
			fmt.Println("stopping process", processName)
		}()

		// set and initialize backoff
		backoffDuration := r.backoffDuration
		backoffMaxRetries := r.backoffMaxRetries
		backoffRetryCount := uint64(0)

		// log starting process
		fmt.Println("starting process", processName)

		for {
			// set process start time
			processStart := time.Now()

			// execute process function
			err := processFunc(r.ctx, r.client, chainId)

			// set process duration
			processDuration := time.Since(processStart)

			// log process duration
			fmt.Println("process duration", processName, processDuration.String())

			if err != nil {
				// log process error
				fmt.Println("process error", processName, err.Error())

				// update retry count
				backoffRetryCount++
			}

			// exit on max retries
			if backoffRetryCount == backoffMaxRetries {
				// log maximum retries
				fmt.Println("maximum retries", processName, backoffMaxRetries)

				// exit process
				return
			} else if backoffRetryCount > 0 {
				// log retry count
				fmt.Println("retry count", processName, backoffRetryCount)
			}

			// wait for the backoff duration to continue or context done to exit
			select {
			case <-time.After(backoffDuration):
				// log backing off
				fmt.Println("backing off", processName, backoffDuration.String())
				// continue process
				continue
			case <-r.ctx.Done():
				// exit process
				return
			}
		}
	}()
}

// Close closes the runner.
func (r *Runner) Close() {
	// log finishing processes
	fmt.Println("finishing processes")

	// wait for processes to finish
	r.waitGroup.Wait()

	// log processes complete
	fmt.Println("processes complete")

	// close indexer client
	err := r.client.Close()
	if err != nil {
		panic(err)
	}

	// log shutdown complete
	fmt.Println("shutdown complete")
}
