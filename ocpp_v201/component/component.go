package component

import "github.com/ChargePi/ocpp-manager/ocpp_v201/variables"

type Component interface {

	// GetName Essentially a component type.
	GetName() Name

	// GetInstanceId returns the unique instance ID of this component.
	GetInstanceId() string

	// RegisterSubComponent registers a sub-component to this component.
	RegisterSubComponent(component Component)

	// UnregisterSubComponent unregisters a sub-component from this component.
	UnregisterSubComponent(component Component)

	// GetSubComponents returns all sub-components of this component.
	GetSubComponents() []Component

	// GetRequiredVariables returns required variables for this component
	GetRequiredVariables() []variables.VariableName

	// GetSupportedVariables returns supported variables (both required and optional) for this component
	GetSupportedVariables() []variables.VariableName

	// GetVariable retrieves a variable by its name.
	GetVariable(key variables.VariableName, opts ...GetSetVariableOption) (*variables.Variable, error)

	// UpdateVariable updates a variable's attribute with a new value.
	UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...GetSetVariableOption) error

	// Validate checks if the variable is valid for this component.
	Validate(key variables.VariableName) bool
}

type Name string

const (
	ComponentNameOCPPCommCtrlr       Name = "OCPPCommCtrlr"
	ComponentNameLocalAuthListCtrlr  Name = "LocalAuthListCtrlr"
	ComponentNameTxCtrlr             Name = "TxCtrlr"
	ComponentNameDeviceDataCtrlr     Name = "DeviceDataCtrlr"
	ComponentNameSecurityCtrlr       Name = "SecurityCtrlr"
	ComponentNameClockCtrlr          Name = "ClockCtrlr"
	ComponentNameCustomizationCtrlr  Name = "CustomizationCtrlr"
	ComponentNameSampledDataCtrlr    Name = "SampledDataCtrlr"
	ComponentNameAlignedDataCtrlr    Name = "AlignedDataCtrlr"
	ComponentNameReservationCtrlr    Name = "ReservationCtrlr"
	ComponentNameSmartChargingCtrlr  Name = "SmartChargingCtrlr"
	ComponentNameTariffCostCtrlr     Name = "TariffCostCtrlr"
	ComponentNameMonitoringCtrlr     Name = "MonitoringCtrlr"
	ComponentNameDisplayMessageCtrlr Name = "DisplayMessageCtrlr"
	ComponentNameISO15118Ctrlr       Name = "ISO15118Ctrlr"
	ComponentNameAuthCtrlr           Name = "AuthCtrlr"
	ComponentNameAuthCacheCtrlr      Name = "AuthCacheCtrlr"
	ComponentNameChargingStation     Name = "ChargingStation"
	ComponentNameEVSE                Name = "EVSE"
	ComponentNameConnector           Name = "Connector"
	ComponentNameConnectedEV         Name = "ConnectedEV"
)
