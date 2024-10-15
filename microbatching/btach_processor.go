package microbatching

import (
	inf "github.com/bimalkeeth/upguard/microbatching/interfaces"
	"time"
)

// startBatchProcessor processes jobs from the jobChannel
// in batches, executing them in the configured size or timeout.
func (mb *MicroBatched[T, R]) startBatchProcessor() {
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

// executeBatch processes a batch of jobs and sends the results through the resultChannel.
func (mb *MicroBatched[T, R]) executeBatch(jobsContainer []inf.Job[T, R]) {
	go func() {
		results := mb.batchProcessor.ProcessBatch(jobsContainer)
		for _, result := range results {
			mb.resultChannel <- result
		}
	}()
}

// Submit adds a job to the micro-batching system.
// If the system is shutting down, it will not accept new jobs.
func (mb *MicroBatched[T, R]) Submit(job inf.Job[T, R]) {
	select {
	case mb.jobChannel <- job:
	case <-mb.shutDownChan:
		return
	}
}

// Start begins the micro-batching process by launching a goroutine
// that listens for incoming jobs and processes them in batches.
func (mb *MicroBatched[T, R]) Start() {
	go func(mb *MicroBatched[T, R]) {
		mb.startBatchProcessor()
	}(mb)
}

// AssignTimer initializes a new timer for the batch processing
// timeout duration using the configured BatchConfig.
func (mb *MicroBatched[T, R]) AssignTimer() {
	mb.batchTimer = time.NewTimer(mb.batchConfig.BatchTimeOutDuration)
}
