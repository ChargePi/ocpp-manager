package ocpp_v201

import (
	"sync"
	"testing"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/controllers"
	"github.com/stretchr/testify/suite"
)

type ManagerTestSuite struct {
	suite.Suite
	manager *Manager
}

func (suite *ManagerTestSuite) SetupTest() {
	suite.manager = NewManager(WithComponents([]component.Component{
		controllers.NewDeviceDataCtrlr(),
	}))
}

func (suite *ManagerTestSuite) TestConcurrentRegisterComponent() {
	manager := NewManager(WithComponents([]component.Component{}))

	var wg sync.WaitGroup
	const goroutines = 20

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			_ = manager.RegisterComponent(controllers.NewClockCtrlr())
		}()
	}
	wg.Wait()

	// Exactly one registration should have succeeded — component must be present.
	_, err := manager.GetVariable(controllers.ComponentNameClockCtrlr, controllers.VariableNameDateTime)
	suite.NotContains(err.Error(), "controller not found")
}

func (suite *ManagerTestSuite) TestConcurrentGetVariable() {
	var wg sync.WaitGroup
	const goroutines = 50

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			v, err := suite.manager.GetVariable(controllers.ComponentNameDeviceDataCtrlr, controllers.VariableNameBytesPerMessage)
			suite.NoError(err)
			suite.NotNil(v)
		}()
	}
	wg.Wait()
}

func (suite *ManagerTestSuite) TestConcurrentUpdateVariable() {
	var wg sync.WaitGroup
	const goroutines = 50

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			// Attribute may not be set; error is expected but must not race.
			_ = suite.manager.UpdateVariable(
				controllers.ComponentNameDeviceDataCtrlr,
				controllers.VariableNameBytesPerMessage,
				"Actual",
				int64(i),
			)
		}(i)
	}
	wg.Wait()
}

func (suite *ManagerTestSuite) TestConcurrentMixedAccess() {
	var wg sync.WaitGroup
	const goroutines = 30

	// Concurrent readers
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			_, _ = suite.manager.GetVariable(controllers.ComponentNameDeviceDataCtrlr, controllers.VariableNameBytesPerMessage)
		}()
	}

	// Concurrent writers
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			_ = suite.manager.UpdateVariable(
				controllers.ComponentNameDeviceDataCtrlr,
				controllers.VariableNameBytesPerMessage,
				"Actual",
				int64(i),
			)
		}(i)
	}

	// Concurrent registrations (all but one will fail with "already registered")
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			_ = suite.manager.RegisterComponent(controllers.NewClockCtrlr())
		}()
	}

	wg.Wait()
}

func (suite *ManagerTestSuite) TestGetVariableControllerNotFound() {
	_, err := suite.manager.GetVariable("NonExistentCtrlr", "SomeVariable")
	suite.EqualError(err, "controller not found")
}

func (suite *ManagerTestSuite) TestUpdateVariableControllerNotFound() {
	err := suite.manager.UpdateVariable("NonExistentCtrlr", "SomeVariable", "Actual", "value")
	suite.EqualError(err, "controller not found")
}

func (suite *ManagerTestSuite) TestRegisterComponentDuplicate() {
	manager := NewManager(WithComponents([]component.Component{}))
	err := manager.RegisterComponent(controllers.NewClockCtrlr())
	suite.NoError(err)

	err = manager.RegisterComponent(controllers.NewClockCtrlr())
	suite.EqualError(err, "component already registered")
}

func TestManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ManagerTestSuite))
}