package worker

import (
	"fmt"
	cons "github.com/bimalkeeth/upguard/microbatching/constants"
	inf "github.com/bimalkeeth/upguard/microbatching/interfaces"
	"github.com/stretchr/testify/mock"
	"time"
)

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Wrong_Batch_Size_Should_Return_Error() {

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            0,
		BatchTimeOutDuration: 0,
	}, suite.mockBatchProcessor)

	err := microBatched.Submit(StrJob[string, string]{
		Id:   1,
		Data: "test",
	})

	suite.Error(err)
	suite.Equal(err.Error(), cons.ErrInvalidBatchSize.Error())
}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Wrong_Duration_Should_Return_Error() {

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            2,
		BatchTimeOutDuration: 0,
	}, suite.mockBatchProcessor)

	err := microBatched.Submit(StrJob[string, string]{
		Id:   1,
		Data: "test",
	})

	suite.Error(err)
	suite.Equal(err.Error(), cons.ErrTimeDuration.Error())
}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Empty_Or_Nill_Job_Should_Return_Error() {

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            2,
		BatchTimeOutDuration: 3,
	}, suite.mockBatchProcessor)

	err := microBatched.Submit(nil)

	suite.Error(err)
	suite.Equal(err.Error(), cons.ErrJobCannotBeNil.Error())
}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Success() {
	var rres inf.JobResult[string] = StrJobResult[string]{

		JobId:   1,
		Result:  "test",
		Error:   nil,
		JobName: "test",
		Success: true,
	}

	suite.mockBatchProcessor.On("ProcessBatch", mock.Anything, mock.Anything).Return([]inf.JobResult[string]{rres})

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            1,
		BatchTimeOutDuration: 1,
	}, suite.mockBatchProcessor)

	err := microBatched.Submit(StrJob[string, string]{
		Id:   1,
		Data: "test",
	})

	suite.NoError(err)
	go func() {
		microBatched.Shutdown()
	}()
}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Success_With_More_Jobs() {
	var rres inf.JobResult[string] = StrJobResult[string]{
		JobId:   1,
		Result:  "test",
		Error:   nil,
		JobName: "test",
		Success: true,
	}

	suite.mockBatchProcessor.On("ProcessBatch", mock.Anything, mock.Anything).Return([]inf.JobResult[string]{rres})

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            10,
		BatchTimeOutDuration: 1,
	}, suite.mockBatchProcessor)

	for i := 0; i < 10; i++ {
		err := microBatched.Submit(StrJob[string, string]{
			Id:   i,
			Data: "test",
		})

		suite.NoError(err)
	}

	time.Sleep(10 * time.Second)

	go func() {
		microBatched.Shutdown()
	}()
}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Success_With_Read_Results() {
	var rres inf.JobResult[string] = StrJobResult[string]{
		JobId:   1,
		Result:  "test",
		Error:   nil,
		JobName: "test",
		Success: true,
	}

	suite.mockBatchProcessor.On("ProcessBatch", mock.Anything, mock.Anything).Return([]inf.JobResult[string]{rres})

	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            10,
		BatchTimeOutDuration: 1,
	}, suite.mockBatchProcessor)

	for i := 0; i < 10; i++ {
		err := microBatched.Submit(StrJob[string, string]{
			Id:   i,
			Data: "test",
		})

		suite.NoError(err)
	}

	time.Sleep(10 * time.Second)

	go func() {
		microBatched.Shutdown()
	}()

	results := microBatched.ReadResult()
	for result := range results {
		fmt.Print(result.IsSuccess())
	}
}
