package worker

import (
	"context"
	"github.com/bimalkeeth/upguard/microbatching/interfaces"
	mockinf "github.com/bimalkeeth/upguard/microbatching/mocks/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type WorkerSuiteTest struct {
	suite.Suite
	asserts            *assert.Assertions
	mockBatchProcessor *mockinf.MockBatchProcessor[string, string]
	mockJobs           *mockinf.MockJob[string, string]
	mockJobResult      *mockinf.MockJobResult[string]
	mockMicroBatcher   *mockinf.MockMicroBatcher[string, string]
}

func (suite *WorkerSuiteTest) SetupTest() {
	suite.mockJobs = mockinf.NewMockJob[string, string](suite.T())
	suite.mockBatchProcessor = mockinf.NewMockBatchProcessor[string, string](suite.T())
	suite.mockJobResult = mockinf.NewMockJobResult[string](suite.T())
	suite.mockMicroBatcher = mockinf.NewMockMicroBatcher[string, string](suite.T())
	suite.asserts = assert.New(suite.T())
}

func (suite *WorkerSuiteTest) GetNewMicroBatched(batchConfig BatchConfig) *microBatched[string, string] {
	return &microBatched[string, string]{
		batchConfig:    batchConfig,
		batchTimer:     time.NewTimer(10),
		resultChannel:  make(chan interfaces.JobResult[string]),
		waitGroup:      sync.WaitGroup{},
		batchProcessor: suite.mockBatchProcessor,
		shutDownChan:   make(chan struct{}),
		jobChannel:     make(chan interfaces.Job[string, string]),
	}
}

func (suite *WorkerSuiteTest) GetServiceContext() context.Context {
	return context.Background()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WorkerSuiteTest))
}
