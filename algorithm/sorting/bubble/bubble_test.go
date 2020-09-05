package bubble

import (
	"testing"

	"github.com/c2fo/testify/suite"
)

type TestBubbleSuite struct {
	suite.Suite
}

func NewCoreTestSuite(t *testing.T) {
	suite.Run(t, new(TestBubbleSuite))
}
