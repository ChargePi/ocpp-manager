package controllers

import (
	"testing"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
	"github.com/stretchr/testify/suite"
)

type OCPPCommCtrlrTestSuite struct {
	suite.Suite
	ctrlr *OCPPCommCtrlr
}

func (suite *OCPPCommCtrlrTestSuite) SetupTest() {
	suite.ctrlr = NewOCPPCommCtrlr()
}

func (suite *OCPPCommCtrlrTestSuite) TestNewOCPPCommCtrlr() {
	ctrlr := NewOCPPCommCtrlr()

	suite.NotNil(ctrlr)
	suite.Equal(component.ComponentNameOCPPCommCtrlr, ctrlr.GetName())
	suite.Equal("ocpp-comm-ctrlr-1", ctrlr.GetInstanceId())
	suite.NotNil(ctrlr.validator)
	suite.NotEmpty(ctrlr.GetRequiredVariables())
	suite.NotEmpty(ctrlr.GetSupportedVariables())
}

func (suite *OCPPCommCtrlrTestSuite) TestGetName() {
	suite.Equal(component.ComponentNameOCPPCommCtrlr, suite.ctrlr.GetName())
}

func (suite *OCPPCommCtrlrTestSuite) TestGetInstanceId() {
	suite.Equal("ocpp-comm-ctrlr-1", suite.ctrlr.GetInstanceId())
}

func (suite *OCPPCommCtrlrTestSuite) TestSubComponentMethods() {
	// Test that sub-component methods are no-op
	initialCount := len(suite.ctrlr.GetSubComponents())

	// Register should be no-op
	suite.ctrlr.RegisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "RegisterSubComponent should be a no-op")

	// Unregister should be no-op
	suite.ctrlr.UnregisterSubComponent(nil)
	suite.Equal(initialCount, len(suite.ctrlr.GetSubComponents()), "UnregisterSubComponent should be a no-op")
}

func (suite *OCPPCommCtrlrTestSuite) TestGetSubComponents() {
	subComponents := suite.ctrlr.GetSubComponents()
	suite.Equal(0, len(subComponents))
}

func (suite *OCPPCommCtrlrTestSuite) TestGetRequiredVariables() {
	requiredVars := suite.ctrlr.GetRequiredVariables()
	suite.NotEmpty(requiredVars)

	expectedVars := []variables.VariableName{
		VariableNameDefaultMessageTimeout,
		VariableNameFileTransferProtocols,
		VariableNameOfflineThreshold,
		VariableNameNetworkProfileConnectionAttempts,
		VariableNameMessageAttempts,
		VariableNameMessageAttemptInterval,
		VariableNameMessageAttemptsTransactionEvent,
		VariableNameMessageAttemptIntervalTransactionEvent,
		VariableNameUnlockOnEVSideDisconnect,
		VariableNameResetRetries,
	}

	for _, expectedVar := range expectedVars {
		suite.Contains(requiredVars, expectedVar)
	}
}

func (suite *OCPPCommCtrlrTestSuite) TestGetSupportedVariables() {
	supportedVars := suite.ctrlr.GetSupportedVariables()
	suite.NotEmpty(supportedVars)

	// Should include all required variables
	requiredVars := suite.ctrlr.GetRequiredVariables()
	for _, requiredVar := range requiredVars {
		suite.Contains(supportedVars, requiredVar)
	}

	// Should include optional variables
	optionalVars := []variables.VariableName{
		VariableNameNetworkConfigurationPriority,
		VariableNameHeartbeatInterval,
		VariableNameMessageTimeout,
		VariableNameActiveNetworkProfile,
		VariableNameQueueAllMessages,
		VariableNameWebSocketPingInterval,
		VariableNameFieldLength,
		VariableNamePublicKeyWithSignedMeterValue,
	}

	for _, optionalVar := range optionalVars {
		suite.Contains(supportedVars, optionalVar)
	}
}

func (suite *OCPPCommCtrlrTestSuite) TestGetVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameHeartbeatInterval,
		variables.VariableTypeInteger,
		int64(300),
	)
	suite.ctrlr.variables[VariableNameHeartbeatInterval] = *testVar

	tests := []struct {
		name          string
		variableName  variables.VariableName
		expectError   bool
		errorContains string
	}{
		{
			name:         "existing variable",
			variableName: VariableNameHeartbeatInterval,
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

func (suite *OCPPCommCtrlrTestSuite) TestUpdateVariable() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameHeartbeatInterval,
		variables.VariableTypeInteger,
		int64(300),
	)
	suite.ctrlr.variables[VariableNameHeartbeatInterval] = *testVar

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
			variableName: VariableNameHeartbeatInterval,
			attribute:    "value",
			value:        int64(600),
			expectError:  false,
		},
		{
			name:          "update non-existent variable",
			variableName:  "NonExistentVariable",
			attribute:     "value",
			value:         int64(600),
			expectError:   true,
			errorContains: "not found",
		},
		{
			name:          "update unsupported variable",
			variableName:  "UnsupportedVariable",
			attribute:     "value",
			value:         int64(600),
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

func (suite *OCPPCommCtrlrTestSuite) TestValidate() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameHeartbeatInterval,
		variables.VariableTypeInteger,
		int64(300),
	)
	suite.ctrlr.variables[VariableNameHeartbeatInterval] = *testVar

	tests := []struct {
		name         string
		variableName variables.VariableName
		expected     bool
	}{
		{
			name:         "valid existing variable",
			variableName: VariableNameHeartbeatInterval,
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

func (suite *OCPPCommCtrlrTestSuite) TestThreadSafety() {
	// Setup a test variable
	testVar := variables.NewVariable(
		VariableNameHeartbeatInterval,
		variables.VariableTypeInteger,
		int64(300),
	)
	suite.ctrlr.variables[VariableNameHeartbeatInterval] = *testVar

	// Test concurrent access to the controller
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			// Concurrent reads
			_, _ = suite.ctrlr.GetVariable(VariableNameHeartbeatInterval)
			_ = suite.ctrlr.Validate(VariableNameHeartbeatInterval)
			_ = suite.ctrlr.GetRequiredVariables()
			_ = suite.ctrlr.GetSupportedVariables()

			// Concurrent writes
			_ = suite.ctrlr.UpdateVariable(VariableNameHeartbeatInterval, "value", int64(600))

			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify the controller is still in a valid state
	suite.True(suite.ctrlr.Validate(VariableNameHeartbeatInterval))
}

func TestOCPPCommCtrlrTestSuite(t *testing.T) {
	suite.Run(t, new(OCPPCommCtrlrTestSuite))
}
