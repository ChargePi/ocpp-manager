package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameCentralContractValidationAllowed               variables.VariableName = "CentralContractValidationAllowed"
	VariableNameContractValidationOffline                      variables.VariableName = "ContractValidationOffline"
	VariableNameProtocolSupportedByEV                          variables.VariableName = "ProtocolSupportedByEV"
	VariableNameProtocolAgreed                                 variables.VariableName = "ProtocolAgreed"
	VariableNameISO15118PnCEnabled                             variables.VariableName = "PnCEnabled"
	VariableNameISO15118V2GCertificateInstallationEnabled      variables.VariableName = "V2GCertificateInstallationEnabled"
	VariableNameISO15118ContractCertificateInstallationEnabled variables.VariableName = "ContractCertificateInstallationEnabled"
	VariableNameISO15118RequestMeteringReceipt                 variables.VariableName = "RequestMeteringReceipt"
	VariableNameISO15118SeccId                                 variables.VariableName = "SeccId"
	VariableNameISO15118CountryName                            variables.VariableName = "CountryName"
	VariableNameISO15118EvseId                                 variables.VariableName = "ISO15118EvseId"
)

func requiredVariablesISO15118Ctrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameCentralContractValidationAllowed,
		VariableNameContractValidationOffline,
		VariableNameProtocolSupportedByEV,
		VariableNameProtocolAgreed,
	}
}

func optionalVariablesISO15118Ctrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameISO15118PnCEnabled,
		VariableNameISO15118V2GCertificateInstallationEnabled,
		VariableNameISO15118ContractCertificateInstallationEnabled,
		VariableNameISO15118RequestMeteringReceipt,
		VariableNameISO15118SeccId,
		VariableNameISO15118CountryName,
		VariableNameISO15118EvseId,
	}
}

func supportedVariablesISO15118Ctrlr() []variables.VariableName {
	return append(requiredVariablesISO15118Ctrlr(), optionalVariablesISO15118Ctrlr()...)
}

type ISO15118Ctrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (I *ISO15118Ctrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (I *ISO15118Ctrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewISO15118Ctrlr() *ISO15118Ctrlr {
	return &ISO15118Ctrlr{
		variables:          map[variables.VariableName]variables.Variable{},
		requiredVariables:  requiredVariablesISO15118Ctrlr(),
		supportedVariables: supportedVariablesISO15118Ctrlr(),
	}
}
