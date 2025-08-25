package controllers

import (
	"testing"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"github.com/stretchr/testify/suite"
)

type TxCtrlrTestSuite struct {
	suite.Suite
	ctrlr *TxCtrlr
}

func (suite *TxCtrlrTestSuite) SetupTest() {
	suite.ctrlr = NewTxCtrlr()
}

func (suite *TxCtrlrTestSuite) TestNewTxCtrlr() {
	ctrlr := NewTxCtrlr()

	suite.NotNil(ctrlr)
	suite.Equal(component.ComponentNameTxCtrlr, ctrlr.GetName())
	suite.Equal("tx-ctrlr", ctrlr.GetInstanceId())
	suite.NotEmpty(ctrlr.GetRequiredVariables())
	suite.NotEmpty(ctrlr.GetSupportedVariables())
}

func (suite *TxCtrlrTestSuite) TestGetName() {
	suite.Equal(component.ComponentNameTxCtrlr, suite.ctrlr.GetName())
}

func (suite *TxCtrlrTestSuite) TestGetInstanceId() {
	suite.Equal("tx-ctrlr", suite.ctrlr.GetInstanceId())
}

func (suite *TxCtrlrTestSuite) TestSubComponentMethods() {
	// Test that sub-component methods are no-op
	initialCount := len(suite.ctrlr.GetSubComponents())

	// Register should be no-op
	suite.ctrlr.RegisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "RegisterSubComponent should be a no-op")

	// Unregister should be no-op
	suite.ctrlr.UnregisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "UnregisterSubComponent should be a no-op")
}

func (suite *TxCtrlrTestSuite) TestGetSubComponents() {
	subComponents := suite.ctrlr.GetSubComponents()
	suite.Equal(0, len(subComponents))
}

func (suite *TxCtrlrTestSuite) TestGetRequiredVariables() {
	requiredVars := suite.ctrlr.GetRequiredVariables()
	expectedVars := requiredTxCtrlrVariables()
	suite.NotEmpty(requiredVars)
	for _, expectedVar := range expectedVars {
		suite.Contains(requiredVars, expectedVar)
	}
}

func (suite *TxCtrlrTestSuite) TestGetSupportedVariables() {
	supportedVars := suite.ctrlr.GetSupportedVariables()
	suite.NotEmpty(supportedVars)

	// Should include all required variables
	requiredVars := suite.ctrlr.GetRequiredVariables()
	for _, requiredVar := range requiredVars {
		suite.Contains(supportedVars, requiredVar)
	}
}

func (suite *TxCtrlrTestSuite) TestGetVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameTxStartPoint,
		variables.VariableTypeString,
		"EVConnected",
	)
	suite.ctrlr.variables[VariableNameTxStartPoint] = *testVar

	tests := []struct {
		name          string
		variableName  variables.VariableName
		expectError   bool
		errorContains string
	}{
		{
			name:         "existing variable",
			variableName: VariableNameTxStartPoint,
			expectError:  false,
		},
		{
			name:          "non-existent variable",
			variableName:  VariableNameDateTime,
			expectError:   true,
			errorContains: "variable DateTime is not supported by this controller",
		},
		{
			name:          "unsupported variable",
			variableName:  "UnsupportedVariable",
			expectError:   true,
			errorContains: "variable UnsupportedVariable is not supported",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			result, err := suite.ctrlr.GetVariable(tt.variableName)

			if tt.expectError {
				suite.Error(err)
				if tt.errorContains != "" {
					suite.Contains(err.Error(), tt.errorContains)
				}
				suite.Nil(result)
			} else {
				suite.NoError(err)
				suite.NotNil(result)
				suite.Equal(tt.variableName, result.Name)
			}
		})
	}
}

func (suite *TxCtrlrTestSuite) TestUpdateVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameTxStartPoint,
		variables.VariableTypeString,
		"EVConnected",
	)
	suite.ctrlr.variables[VariableNameTxStartPoint] = *testVar

	tests := []struct {
		name          string
		variableName  variables.VariableName
		attribute     string
		value         interface{}
		expectError   bool
		errorContains string
	}{
		{
			name:         "update existing variable",
			variableName: VariableNameTxStartPoint,
			attribute:    "value",
			value:        "Authorized",
			expectError:  false,
		},
		{
			name:          "update non-existent variable",
			variableName:  "NonExistentVariable",
			attribute:     "value",
			value:         "Authorized",
			expectError:   true,
			errorContains: "variable NonExistentVariable is not supported by this controller",
		},
		{
			name:          "update unsupported variable",
			variableName:  "UnsupportedVariable",
			attribute:     "value",
			value:         "Authorized",
			expectError:   true,
			errorContains: "not supported",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.ctrlr.UpdateVariable(tt.variableName, tt.attribute, tt.value)

			if tt.expectError {
				suite.Error(err)
				if tt.errorContains != "" {
					suite.Contains(err.Error(), tt.errorContains)
				}
			} else {
				suite.NoError(err)

				// Verify the update worked
				variable, err := suite.ctrlr.GetVariable(tt.variableName)
				suite.NoError(err)
				suite.NotNil(variable)
			}
		})
	}
}

func (suite *TxCtrlrTestSuite) TestValidate() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameTxStartPoint,
		variables.VariableTypeString,
		"EVConnected",
	)
	suite.ctrlr.variables[VariableNameTxStartPoint] = *testVar

	tests := []struct {
		name         string
		variableName variables.VariableName
		expected     bool
	}{
		{
			name:         "valid existing variable",
			variableName: VariableNameTxStartPoint,
			expected:     true,
		},
		{
			name:         "non-existent variable",
			variableName: "NonExistentVariable",
			expected:     false,
		},
		{
			name:         "unsupported variable",
			variableName: "UnsupportedVariable",
			expected:     false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			result := suite.ctrlr.Validate(tt.variableName)
			suite.Equal(tt.expected, result)
		})
	}
}

func (suite *TxCtrlrTestSuite) TestThreadSafety() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameTxStartPoint,
		variables.VariableTypeString,
		"EVConnected",
	)
	suite.ctrlr.variables[VariableNameTxStartPoint] = *testVar

	// Test concurrent access to the controller
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			// Concurrent reads
			_, _ = suite.ctrlr.GetVariable(VariableNameTxStartPoint)
			_ = suite.ctrlr.Validate(VariableNameTxStartPoint)
			_ = suite.ctrlr.GetRequiredVariables()
			_ = suite.ctrlr.GetSupportedVariables()

			// Concurrent writes
			_ = suite.ctrlr.UpdateVariable(VariableNameTxStartPoint, "value", "Authorized")

			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify the controller is still in a valid state
	suite.True(suite.ctrlr.Validate(VariableNameTxStartPoint))
}

func TestTxCtrlrTestSuite(t *testing.T) {
	suite.Run(t, new(TxCtrlrTestSuite))
}
