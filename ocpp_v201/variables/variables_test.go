package variables

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type variablesTestSuite struct {
	suite.Suite
}

func (suite *variablesTestSuite) TestVariable_Validate() {

}

func (suite *variablesTestSuite) TestVariable_GetVariableAttribute() {

}

func (suite *variablesTestSuite) TestVariable_UpdateVariableAttribute() {

}

func (suite *variablesTestSuite) TestVariableAttributes_Validate() {

}

func (suite *variablesTestSuite) TestVariableAttributes_Update() {

}

func TestVariables(t *testing.T) {
	suite.Run(t, new(variablesTestSuite))
}
