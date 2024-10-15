package worker

import (
	cons "github.com/bimalkeeth/upguard/microbatching/constants"
	inf "github.com/bimalkeeth/upguard/microbatching/interfaces"
	"time"
)

// startBatchProcessor processes jobs from the jobChannel
// in batches, executing them in the configured size or timeout.
func (mb *microBatched[T, R]) startBatchProcessor() {
	defer mb.waitGroup.Done()
	var jobsContainer []inf.Job[T, R]

	for {
		select {
		case job := <-mb.jobChannel:
			jobsContainer = append(jobsContainer, job)
			if len(jobsContainer) >= mb.batchConfig.BatchSize {
				mb.executeBatch(jobsContainer)
				jobsContainer = []inf.Job[T, R]{}

				mb.batchTimer.Reset(mb.batchConfig.BatchTimeOutDuration)
			}
		case <-mb.batchTimer.C:
			if len(jobsContainer) > 0 {
				mb.executeBatch(jobsContainer)
			}
			return
		case <-mb.shutDownChan:
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
		results := mb.batchProcessor.ProcessBatch(jobsContainer)
		for _, result := range results {
			mb.resultChannel <- result
		}
	}()
}

// Submit adds a job to the micro-batching system.
// If the system is shutting down, it will not accept new jobs.
func (mb *microBatched[T, R]) Submit(job inf.Job[T, R]) error {
	if job == nil {
		return cons.ErrJobCannotBeNil
	}

	if mb.batchConfig.BatchSize == 0 {
		return cons.ErrInvalidBatchSize
	}

	if mb.batchConfig.BatchTimeOutDuration.Seconds() == 0 {
		return cons.ErrTimeDuration
	}

	select {
	case mb.jobChannel <- job:
	case <-mb.shutDownChan:
		return nil
	}

	return nil
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
