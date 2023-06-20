package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

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

	fmt.Println("runner", "shutdown complete")
}

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(f process.Function, p process.Params) {
	// validate process parameters
	err := p.ValidateParams()
	if err != nil {
		defer fmt.Println("runner", "invalid params", p.Name)
		return // exit process
	}

	// initialize new client if not empty
	if p.ChainRpc != "" {
		fmt.Println("runner", "initializing new client", p.ChainId, p.ChainRpc)

		// TODO: initialize new client
		fmt.Println("runner", "error", "not implemented", p.Name)

		return // exit process
	}

	// add process to wait group
	r.waitGroup.Add(1)

	go func() {
		// decrement wait group counter on exit
		defer r.waitGroup.Done()

		defer fmt.Println("runner", "stopping process", p.Name)

		// set and initialize backoff
		backoffDuration := r.cfg.RunnerBackoffDuration
		backoffMaxRetries := r.cfg.RunnerBackoffMaxRetries
		backoffRetryCount := uint64(0)

		fmt.Println("runner", "starting process", p.Name)

		for {
			// log retry count
			if backoffRetryCount > 0 {
				fmt.Println("runner", "retry count", p.Name, backoffRetryCount)
			}

			// exit on exceeding max retries
			if backoffRetryCount > backoffMaxRetries {
				fmt.Println("runner", "maximum retries", p.Name, backoffMaxRetries)
				return // exit process
			}

			// set process start time
			processStart := time.Now()

			// execute process function
			err := f(r.ctx, p)

			// set process duration
			processDuration := time.Since(processStart)

			fmt.Println("runner", "process duration", p.Name, processDuration.String())

			if err != nil {
				fmt.Println("runner", "process error", p.Name, err.Error())

				// update retry count
				backoffRetryCount++
			}

			// wait for context done or backoff duration
			select {
			case <-r.ctx.Done():
				return // exit process
			case <-time.After(backoffDuration):
				fmt.Println("runner", "backing off", p.Name, backoffDuration.String())
				continue // continue process
			}
		}
	}()
}
