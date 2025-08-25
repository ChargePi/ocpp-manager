package ocpp_v201

import "github.com/ChargePi/ocpp-manager/ocpp_v201/component"

type Connector struct {
	ID         int
	components map[component.Name]component.Component
}
