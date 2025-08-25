package controllers

import (
	"testing"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"github.com/stretchr/testify/suite"
)

type DeviceDataCtrlrTestSuite struct {
	suite.Suite
	ctrlr *DeviceDataCtrlr
}

func (suite *DeviceDataCtrlrTestSuite) SetupTest() {
	suite.ctrlr = NewDeviceDataCtrlr()
}

func (suite *DeviceDataCtrlrTestSuite) TestNewDeviceDataCtrlr() {
	ctrlr := NewDeviceDataCtrlr()

	suite.NotNil(ctrlr)
	suite.Equal(component.ComponentNameDeviceDataCtrlr, ctrlr.GetName())
	suite.Equal("device-data-ctrlr", ctrlr.GetInstanceId())
	suite.NotNil(ctrlr.validator)
	suite.NotEmpty(ctrlr.GetRequiredVariables())
	suite.NotEmpty(ctrlr.GetSupportedVariables())
}

func (suite *DeviceDataCtrlrTestSuite) TestGetName() {
	suite.Equal(component.ComponentNameDeviceDataCtrlr, suite.ctrlr.GetName())
}

func (suite *DeviceDataCtrlrTestSuite) TestGetInstanceId() {
	suite.Equal("device-data-ctrlr", suite.ctrlr.GetInstanceId())
}

func (suite *DeviceDataCtrlrTestSuite) TestSubComponentMethods() {
	// Test that sub-component methods are no-op
	initialCount := len(suite.ctrlr.GetSubComponents())

	// Register should be no-op
	suite.ctrlr.RegisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "RegisterSubComponent should be a no-op")

	// Unregister should be no-op
	suite.ctrlr.UnregisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "UnregisterSubComponent should be a no-op")
}

func (suite *DeviceDataCtrlrTestSuite) TestGetSubComponents() {
	subComponents := suite.ctrlr.GetSubComponents()
	suite.Equal(0, len(subComponents))
}

func (suite *DeviceDataCtrlrTestSuite) TestGetRequiredVariables() {
	requiredVars := suite.ctrlr.GetRequiredVariables()
	suite.NotEmpty(requiredVars)

	expectedVars := []variables.VariableName{
		VariableNameBytesPerMessage,
		VariableNameConfigurationValueSize,
		VariableNameReportingValueSize,
	}

	for _, expectedVar := range expectedVars {
		suite.Contains(requiredVars, expectedVar)
	}
}

func (suite *DeviceDataCtrlrTestSuite) TestGetSupportedVariables() {
	supportedVars := suite.ctrlr.GetSupportedVariables()
	suite.NotEmpty(supportedVars)

	// Should include all required variables
	requiredVars := suite.ctrlr.GetRequiredVariables()
	for _, requiredVar := range requiredVars {
		suite.Contains(supportedVars, requiredVar)
	}
}

func (suite *DeviceDataCtrlrTestSuite) TestGetVariable() {
	tests := []struct {
		name          string
		variableName  variables.VariableName
		setupVariable bool
		expectError   bool
		errorContains string
	}{
		{
			name:          "existing variable",
			variableName:  VariableNameBytesPerMessage,
			setupVariable: true,
			expectError:   false,
		},
		{
			name:          "non-existent variable",
			variableName:  "NonExistentVariable",
			setupVariable: false,
			expectError:   true,
			errorContains: "not found",
		},
		{
			name:          "unsupported variable",
			variableName:  "UnsupportedVariable",
			setupVariable: false,
			expectError:   true,
			errorContains: "not supported",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if tt.setupVariable {
				// Variable is already set up in constructor
			}

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

func (suite *DeviceDataCtrlrTestSuite) TestUpdateVariable() {
	tests := []struct {
		name          string
		variableName  variables.VariableName
		attribute     string
		value         interface{}
		setupVariable bool
		expectError   bool
		errorContains string
	}{
		{
			name:          "update existing variable",
			variableName:  VariableNameBytesPerMessage,
			attribute:     "value",
			value:         int64(1024),
			setupVariable: true,
			expectError:   false,
		},
		{
			name:          "update non-existent variable",
			variableName:  "NonExistentVariable",
			attribute:     "value",
			value:         int64(1024),
			setupVariable: false,
			expectError:   true,
			errorContains: "not found",
		},
		{
			name:          "update unsupported variable",
			variableName:  "UnsupportedVariable",
			attribute:     "value",
			value:         int64(1024),
			setupVariable: false,
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

func (suite *DeviceDataCtrlrTestSuite) TestValidate() {
	tests := []struct {
		name          string
		variableName  variables.VariableName
		setupVariable bool
		expected      bool
	}{
		{
			name:          "valid existing variable",
			variableName:  VariableNameBytesPerMessage,
			setupVariable: true,
			expected:      true,
		},
		{
			name:          "non-existent variable",
			variableName:  "NonExistentVariable",
			setupVariable: false,
			expected:      false,
		},
		{
			name:          "unsupported variable",
			variableName:  "UnsupportedVariable",
			setupVariable: false,
			expected:      false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			result := suite.ctrlr.Validate(tt.variableName)
			suite.Equal(tt.expected, result)
		})
	}
}

func (suite *DeviceDataCtrlrTestSuite) TestValidateAllRequiredVariables() {
	// Test with all required variables present (default state)
	result := suite.ctrlr.ValidateAllRequiredVariables()
	suite.True(result)
}

func (suite *DeviceDataCtrlrTestSuite) TestThreadSafety() {
	// Test concurrent access to the controller
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			// Concurrent reads
			_, _ = suite.ctrlr.GetVariable(VariableNameBytesPerMessage)
			_ = suite.ctrlr.Validate(VariableNameBytesPerMessage)
			_ = suite.ctrlr.GetRequiredVariables()
			_ = suite.ctrlr.GetSupportedVariables()

			// Concurrent writes
			_ = suite.ctrlr.UpdateVariable(VariableNameBytesPerMessage, "value", int64(512))

			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify the controller is still in a valid state
	suite.True(suite.ctrlr.ValidateAllRequiredVariables())
}

func TestDeviceDataCtrlrTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceDataCtrlrTestSuite))
}
