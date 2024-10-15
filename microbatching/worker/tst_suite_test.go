package worker

import (
	mockinf "github.com/bimalkeeth/upguard/microbatching/mocks/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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

func TestSuite(t *testing.T) {
	suite.Run(t, new(WorkerSuiteTest))
}
