package worker

import (
	cons "github.com/bimalkeeth/upguard/microbatching/constants"
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
	microBatched := NewMicroBatched[string, string](suite.GetServiceContext(), BatchConfig{
		BatchSize:            1,
		BatchTimeOutDuration: 1,
	}, suite.mockBatchProcessor)

	err := microBatched.Submit(StrJob[string, string]{
		Id:   1,
		Data: "test",
	})

	suite.NoError(err)

}

func (suite *WorkerSuiteTest) Test_BatchProcessor_Process_Submit_Batch_Success_With_More_Jobs() {
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
}
