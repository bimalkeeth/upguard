package worker

import (
	"sync"
	"time"

	cons "github.com/bimalkeeth/upguard/microbatching/constants"
	inf "github.com/bimalkeeth/upguard/microbatching/interfaces"
)

var shutDownOnce sync.Once
var jobChannelOnce sync.Once
var resultChannelOnce sync.Once

// startBatchProcessor processes jobs from the jobChannel
// in batches, executing them in the configured size or timeout.
func (mb *microBatched[T, R]) startBatchProcessor() {
	defer mb.waitGroup.Done()

	var jobsContainer []inf.Job[T, R]

	for {
		select {
		case job, ok := <-mb.jobChannel:
			if !ok {
				// jobChannel has been closed and drained, process remaining jobs and exit.
				if len(jobsContainer) > 0 {
					mb.executeBatch(jobsContainer)
				}

				return
			}

			jobsContainer = append(jobsContainer, job)
			if len(jobsContainer) >= mb.batchConfig.BatchSize {
				mb.executeBatch(jobsContainer)
				jobsContainer = []inf.Job[T, R]{}

				mb.batchTimer.Reset(mb.batchConfig.BatchTimeOutDuration)
			}
		case <-mb.batchTimer.C:
			if len(jobsContainer) > 0 {
				mb.executeBatch(jobsContainer)
				jobsContainer = []inf.Job[T, R]{}
			}
		case <-mb.shutDownChan:
			// Received shutdown signal, process any remaining jobs
			mb.batchTimer.Stop()

			if len(jobsContainer) > 0 {
				mb.executeBatch(jobsContainer)
			}

			return
		}
	}
}

// executeBatch processes a worker of jobs and sends the results through the resultChannel.
func (mb *microBatched[T, R]) executeBatch(jobsContainer []inf.Job[T, R]) {
	go func() {
		results := mb.batchProcessor.ProcessBatch(mb.serviceCtx, jobsContainer)
		for _, result := range results {
			select {
			case mb.resultChannel <- result:
			case <-mb.shutDownChan:
				// Ensure we are not sending on a closed resultChannel
				return
			}
		}
	}()
}

// Submit adds a job to the micro-batching system.
// If the system is shutting down, it will not accept new jobs.
func (mb *microBatched[T, R]) Submit(job inf.Job[T, R]) error {
	if mb.batchConfig.BatchSize == 0 {
		return cons.ErrInvalidBatchSize
	}

	if mb.batchConfig.BatchTimeOutDuration.Seconds() == 0 {
		return cons.ErrTimeDuration
	}

	if job == nil {
		return cons.ErrJobCannotBeNil
	}

	select {
	case mb.jobChannel <- job:
	case <-mb.shutDownChan:
		// System is shutting down, do not accept new jobs.
		return nil
	}

	return nil
}

// Shutdown gracefully stops the micro-batching process
// And waits for all goroutines to complete.
func (mb *microBatched[T, R]) Shutdown() {
	// Close shutDownChan once
	onceClose(mb.shutDownChan, &shutDownOnce)
	// Close jobChannel once
	onceClose(mb.jobChannel, &jobChannelOnce)
	// Wait for the batch processor to finish processing all jobs
	mb.waitGroup.Wait()
	// Close resultChannel once all jobs are done
	onceClose(mb.resultChannel, &resultChannelOnce)
}

// Start begins the micro-batching worker by launching a goroutine
// that listens for incoming jobs and processes them in batches.
func (mb *microBatched[T, R]) start() {
	go func(mb *microBatched[T, R]) {
		mb.startBatchProcessor()
	}(mb)
}

// AssignTimer initializes a new timer for the worker processing
// timeout duration using the configured BatchConfig.
func (mb *microBatched[T, R]) assignTimer() {
	mb.batchTimer = time.NewTimer(mb.batchConfig.BatchTimeOutDuration)
}

// ReadResult streams job results from the micro-batching system.
// It returns a channel where job results of type inf.JobResult[R] can be received.
func (mb *microBatched[T, R]) ReadResult() <-chan inf.JobResult[R] {
	resultRetChan := make(chan inf.JobResult[R])
	go func(mb *microBatched[T, R], resultRetChan chan inf.JobResult[R]) {
		defer close(resultRetChan)
		for {
			select {
			case result, ok := <-mb.resultChannel:
				if !ok {
					return
				}
				resultRetChan <- result
			case <-mb.shutDownChan:
				return
			}
		}
	}(mb, resultRetChan)

	return resultRetChan
}

func onceClose[T any](ch chan T, once *sync.Once) {
	once.Do(func() {
		close(ch)
	})
}
