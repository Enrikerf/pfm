package ModelTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
)

type outPinMock struct {}

func NewOutPinMock() Pin.OutPin {
	return &outPinMock{}
}

func (outPin *outPinMock) TearDown() {}

func (outPin *outPinMock) Up() {}

func (outPin *outPinMock) Down() {}
