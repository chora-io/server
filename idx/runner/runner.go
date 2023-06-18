package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
)

// Runner runs continuous background processes.
type Runner struct {
	ctx       context.Context
	cfg       config.Config
	client    client.Client
	waitGroup sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, cfg config.Config, client client.Client) Runner {
	return Runner{
		ctx:    ctx,
		cfg:    cfg,
		client: client,
	}
}

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(processName string, processFunc process.Function) {
	// add process to wait group
	r.waitGroup.Add(1)

	go func() {
		// decrement wait group counter on exit
		defer r.waitGroup.Done()

		defer fmt.Println("stopping process", processName)

		// set and initialize backoff
		backoffDuration := r.cfg.BackoffDuration
		backoffMaxRetries := r.cfg.BackoffMaxRetries
		backoffRetryCount := uint64(0)

		fmt.Println("starting process", processName)

		for {
			// log retry count
			if backoffRetryCount > 0 {
				fmt.Println("retry count", processName, backoffRetryCount)
			}

			// exit on exceeding max retries
			if backoffRetryCount > backoffMaxRetries {
				fmt.Println("maximum retries", processName, backoffMaxRetries)
				return // exit process
			}

			// set process start time
			processStart := time.Now()

			// execute process function
			err := processFunc(r.ctx, r.client)

			// set process duration
			processDuration := time.Since(processStart)

			fmt.Println("process duration", processName, processDuration.String())

			if err != nil {
				fmt.Println("process error", processName, err.Error())

				// update retry count
				backoffRetryCount++
			}

			// wait for context done or backoff duration
			select {
			case <-r.ctx.Done():
				return // exit process
			case <-time.After(backoffDuration):
				fmt.Println("backing off", processName, backoffDuration.String())
				continue // continue process
			}
		}
	}()
}

// Close shuts down the runner.
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
