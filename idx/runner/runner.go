package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

	idxcontext "github.com/choraio/server/idx/context"
)

// Runner runs continuous background processes.
type Runner struct {
	// Context is the context used for cancellation and shutdown of the runner.
	Context context.Context

	// IdxContext is the indexer context, which includes the database client.
	IdxContext idxcontext.Context

	// waitGroup tracks running processes and is used to coordinate a shutdown.
	waitGroup sync.WaitGroup
}

// NewRunner creates a new runner.
func NewRunner(ctx context.Context, idxCtx idxcontext.Context) Runner {
	return Runner{
		Context:    ctx,
		IdxContext: idxCtx,
	}
}

// ProcessFunc is the function used to advance the process.
type ProcessFunc func(ctx context.Context, idxCtx idxcontext.Context) error

// RunProcess runs a process using the provided process function.
func (r *Runner) RunProcess(processName string, processFunc ProcessFunc) {
	ctx := r.Context
	if ctx == nil {
		ctx = context.Background()
	}

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
		backoffDuration := r.IdxContext.BackoffDuration
		backoffMaxRetries := r.IdxContext.BackoffMaxRetries
		backoffRetryCount := uint64(0)

		// log starting process
		fmt.Println("starting process", processName)

		for {
			// set process start time
			processStart := time.Now()

			// execute process function
			err := processFunc(ctx, r.IdxContext)

			// set process duration
			processDuration := time.Since(processStart)

			// log process duration
			fmt.Println("process duration", processName, processDuration)

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
			case <-ctx.Done():
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

	// close indexer context
	err := r.IdxContext.Close()
	if err != nil {
		panic(err)
	}

	// log shutdown complete
	fmt.Println("shutdown complete")
}
