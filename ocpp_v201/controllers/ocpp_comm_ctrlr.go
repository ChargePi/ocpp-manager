package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDefaultMessageTimeout                  variables.VariableName = "DefaultMessageTimeout"
	VariableNameNetworkProfileConnectionAttempts       variables.VariableName = "NetworkProfileConnectionAttempts"
	VariableNameNetworkConfigurationPriority           variables.VariableName = "NetworkConfigurationPriority"
	VariableNameHeartbeatInterval                      variables.VariableName = "HeartbeatInterval"
	VariableNameFileTransferProtocols                  variables.VariableName = "FileTransferProtocols"
	VariableNameMessageTimeout                         variables.VariableName = "MessageTimeout"
	VariableNameActiveNetworkProfile                   variables.VariableName = "ActiveNetworkProfile"
	VariableNameOfflineThreshold                       variables.VariableName = "OfflineThreshold"
	VariableNameQueueAllMessages                       variables.VariableName = "QueueAllMessages"
	VariableNameMessageAttempts                        variables.VariableName = "MessageAttempts"
	VariableNameMessageAttemptInterval                 variables.VariableName = "MessageAttemptInterval"
	VariableNameMessageAttemptsTransactionEvent        variables.VariableName = "MessageAttemptsTransactionEvent"
	VariableNameMessageAttemptIntervalTransactionEvent variables.VariableName = "MessageAttemptIntervalTransactionEvent"
	VariableNameUnlockOnEVSideDisconnect               variables.VariableName = "UnlockOnEVSideDisconnect"
	VariableNameWebSocketPingInterval                  variables.VariableName = "WebSocketPingInterval"
	VariableNameResetRetries                           variables.VariableName = "ResetRetries"
	VariableNameFieldLength                            variables.VariableName = "FieldLength"
	VariableNamePublicKeyWithSignedMeterValue          variables.VariableName = "PublicKeyWithSignedMeterValue"
)

func requiredOcppCommCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
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
}

func optionalOcppCommCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameNetworkConfigurationPriority,
		VariableNameHeartbeatInterval,
		VariableNameMessageTimeout,
		VariableNameActiveNetworkProfile,
		VariableNameQueueAllMessages,
		VariableNameWebSocketPingInterval,
		VariableNameFieldLength,
		VariableNamePublicKeyWithSignedMeterValue,
	}
}

func supportedOcppCommCtrlrVariables() []variables.VariableName {
	return append(requiredOcppCommCtrlrVariables(), optionalOcppCommCtrlrVariables()...)
}

type OCPPCommCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (ctrlr *OCPPCommCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (ctrlr *OCPPCommCtrlr) Validate(key variables.VariableName) bool {
	v, exists := ctrlr.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewOCPPCommCtrlr() *OCPPCommCtrlr {
	return &OCPPCommCtrlr{
		variables:          map[variables.VariableName]variables.Variable{},
		requiredVariables:  requiredOcppCommCtrlrVariables(),
		supportedVariables: supportedOcppCommCtrlrVariables(),
	}
}
