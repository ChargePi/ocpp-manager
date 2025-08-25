package controllers

import (
	"testing"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"github.com/stretchr/testify/suite"
)

type DisplayCtrlrTestSuite struct {
	suite.Suite
	ctrlr *DisplayCtrlr
}

func (suite *DisplayCtrlrTestSuite) SetupTest() {
	suite.ctrlr = NewDisplayCtrlr()
}

func (suite *DisplayCtrlrTestSuite) TestNewDisplayCtrlr() {
	ctrlr := NewDisplayCtrlr()

	suite.NotNil(ctrlr)
	suite.Equal(component.ComponentNameDisplayMessageCtrlr, ctrlr.GetName())
	suite.Equal("display-message-ctrlr", ctrlr.GetInstanceId())
	suite.NotNil(ctrlr.validator)
	suite.NotEmpty(ctrlr.GetRequiredVariables())
	suite.NotEmpty(ctrlr.GetSupportedVariables())
}

func (suite *DisplayCtrlrTestSuite) TestGetName() {
	suite.Equal(component.ComponentNameDisplayMessageCtrlr, suite.ctrlr.GetName())
}

func (suite *DisplayCtrlrTestSuite) TestGetInstanceId() {
	suite.Equal("display-message-ctrlr", suite.ctrlr.GetInstanceId())
}

func (suite *DisplayCtrlrTestSuite) TestSubComponentMethods() {
	// Test that sub-component methods are no-op
	initialCount := len(suite.ctrlr.GetSubComponents())

	// Register should be no-op
	suite.ctrlr.RegisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "RegisterSubComponent should be a no-op")

	// Unregister should be no-op
	suite.ctrlr.UnregisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "UnregisterSubComponent should be a no-op")
}

func (suite *DisplayCtrlrTestSuite) TestGetSubComponents() {
	subComponents := suite.ctrlr.GetSubComponents()
	suite.Equal(0, len(subComponents))
}

func (suite *DisplayCtrlrTestSuite) TestGetRequiredVariables() {
	requiredVars := suite.ctrlr.GetRequiredVariables()
	suite.NotEmpty(requiredVars)

	expectedVars := []variables.VariableName{
		VariableNameNumberOfDisplayMessages,
		VariableNameDisplayMessageSupportedFormats,
		VariableNameDisplayMessageSupportedPriorities,
	}

	for _, expectedVar := range expectedVars {
		suite.Contains(requiredVars, expectedVar)
	}
}

func (suite *DisplayCtrlrTestSuite) TestGetSupportedVariables() {
	supportedVars := suite.ctrlr.GetSupportedVariables()
	suite.NotEmpty(supportedVars)

	// Should include all required variables
	requiredVars := suite.ctrlr.GetRequiredVariables()
	for _, requiredVar := range requiredVars {
		suite.Contains(supportedVars, requiredVar)
	}
}

func (suite *DisplayCtrlrTestSuite) TestGetVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameDisplayMessageSupportedFormats,
		variables.VariableTypeString,
		"ASCII",
	)
	suite.ctrlr.variables[VariableNameDisplayMessageSupportedFormats] = *testVar

	tests := []struct {
		name          string
		variableName  variables.VariableName
		expectError   bool
		errorContains string
	}{
		{
			name:         "existing variable",
			variableName: VariableNameDisplayMessageSupportedFormats,
			expectError:  false,
		},
		{
			name:          "non-existent variable",
			variableName:  "NonExistentVariable",
			expectError:   true,
			errorContains: "not found",
		},
		{
			name:          "unsupported variable",
			variableName:  "UnsupportedVariable",
			expectError:   true,
			errorContains: "not supported",
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

func (suite *DisplayCtrlrTestSuite) TestUpdateVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameDisplayMessageSupportedFormats,
		variables.VariableTypeString,
		"ASCII",
	)
	suite.ctrlr.variables[VariableNameDisplayMessageSupportedFormats] = *testVar

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
			variableName: VariableNameDisplayMessageSupportedFormats,
			attribute:    "value",
			value:        "UTF8",
			expectError:  false,
		},
		{
			name:          "update non-existent variable",
			variableName:  "NonExistentVariable",
			attribute:     "value",
			value:         "UTF8",
			expectError:   true,
			errorContains: "not found",
		},
		{
			name:          "update unsupported variable",
			variableName:  "UnsupportedVariable",
			attribute:     "value",
			value:         "UTF8",
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

func (suite *DisplayCtrlrTestSuite) TestValidate() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameDisplayMessageSupportedFormats,
		variables.VariableTypeString,
		"ASCII",
	)
	suite.ctrlr.variables[VariableNameDisplayMessageSupportedFormats] = *testVar

	tests := []struct {
		name         string
		variableName variables.VariableName
		expected     bool
	}{
		{
			name:         "valid existing variable",
			variableName: VariableNameDisplayMessageSupportedFormats,
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

func (suite *DisplayCtrlrTestSuite) TestThreadSafety() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameDisplayMessageSupportedFormats,
		variables.VariableTypeString,
		"ASCII",
	)
	suite.ctrlr.variables[VariableNameDisplayMessageSupportedFormats] = *testVar

	// Test concurrent access to the controller
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			// Concurrent reads
			_, _ = suite.ctrlr.GetVariable(VariableNameDisplayMessageSupportedFormats)
			_ = suite.ctrlr.Validate(VariableNameDisplayMessageSupportedFormats)
			_ = suite.ctrlr.GetRequiredVariables()
			_ = suite.ctrlr.GetSupportedVariables()

			// Concurrent writes
			_ = suite.ctrlr.UpdateVariable(VariableNameDisplayMessageSupportedFormats, "value", "UTF8")

			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify the controller is still in a valid state
	suite.True(suite.ctrlr.Validate(VariableNameDisplayMessageSupportedFormats))
}

func TestDisplayCtrlrTestSuite(t *testing.T) {
	suite.Run(t, new(DisplayCtrlrTestSuite))
}
