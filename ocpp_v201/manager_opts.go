package ocpp_v201

import "github.com/ChargePi/ocpp-manager/ocpp_v201/component"

type ManagerOption func(*managerOptions)

type managerOptions struct {
	// Supported profiles + components
	components []component.Component
}

func WithComponents(components []component.Component) ManagerOption {
	return func(opts *managerOptions) {
		opts.components = components
	}
}
