package ocpp_v201

type ComponentName string

const (
	ComponentNameOCPPCommCtrlr       ComponentName = "OCPPCommCtrlr"
	ComponentNameLocalAuthListCtrlr  ComponentName = "LocalAuthListCtrlr"
	ComponentNameTxCtrlr             ComponentName = "TxCtrlr"
	ComponentNameDeviceDataCtrlr     ComponentName = "DeviceDataCtrlr"
	ComponentNameSecurityCtrlr       ComponentName = "SecurityCtrlr"
	ComponentNameClockCtrlr          ComponentName = "ClockCtrlr"
	ComponentNameCustomizationCtrlr  ComponentName = "CustomizationCtrlr"
	ComponentNameSampledDataCtrlr    ComponentName = "SampledDataCtrlr"
	ComponentNameAlignedDataCtrlr    ComponentName = "AlignedDataCtrlr"
	ComponentNameReservationCtrlr    ComponentName = "ReservationCtrlr"
	ComponentNameSmartChargingCtrlr  ComponentName = "SmartChargingCtrlr"
	ComponentNameTariffCostCtrlr     ComponentName = "TariffCostCtrlr"
	ComponentNameMonitoringCtrlr     ComponentName = "MonitoringCtrlr"
	ComponentNameDisplayMessageCtrlr ComponentName = "DisplayMessageCtrlr"
	ComponentNameISO15118Ctrlr       ComponentName = "ISO15118Ctrlr"
	ComponentNameAuthCtrlr           ComponentName = "AuthCtrlr"
	ComponentNameAuthCacheCtrlr      ComponentName = "AuthCacheCtrlr"
	ComponentNameChargingStation     ComponentName = "ChargingStation"
	ComponentNameEVSE                ComponentName = "EVSE"
	ComponentNameConnector           ComponentName = "Connector"
	ComponentNameConnectedEV         ComponentName = "ConnectedEV"
)

type Controller interface {
	Get() (*Variable, error)
	Update(key VariableName, value interface{}) error
	Validate(key VariableName) bool
}
