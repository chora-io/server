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
	waitGroup sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, cfg config.Config) Runner {
	return Runner{
		ctx: ctx,
		cfg: cfg,
	}
}

// Close shuts down the runner.
func (r *Runner) Close() {
	// wait for processes to finish
	r.waitGroup.Wait()

	fmt.Println("shutdown complete")
}

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(c client.Client, f process.Function, p process.Params) {
	// add process to wait group
	r.waitGroup.Add(1)

	go func() {
		// decrement wait group counter on exit
		defer r.waitGroup.Done()

		defer fmt.Println("stopping process", p.Name)

		// set and initialize backoff
		backoffDuration := r.cfg.BackoffDuration
		backoffMaxRetries := r.cfg.BackoffMaxRetries
		backoffRetryCount := uint64(0)

		fmt.Println("starting process", p.Name)

		for {
			// log retry count
			if backoffRetryCount > 0 {
				fmt.Println("retry count", p.Name, backoffRetryCount)
			}

			// exit on exceeding max retries
			if backoffRetryCount > backoffMaxRetries {
				fmt.Println("maximum retries", p.Name, backoffMaxRetries)
				return // exit process
			}

			// set process start time
			processStart := time.Now()

			// execute process function
			err := f(r.ctx, c, p)

			// set process duration
			processDuration := time.Since(processStart)

			fmt.Println("process duration", p.Name, processDuration.String())

			if err != nil {
				fmt.Println("process error", p.Name, err.Error())

				// update retry count
				backoffRetryCount++
			}

			// wait for context done or backoff duration
			select {
			case <-r.ctx.Done():
				return // exit process
			case <-time.After(backoffDuration):
				fmt.Println("backing off", p.Name, backoffDuration.String())
				continue // continue process
			}
		}
	}()
}
