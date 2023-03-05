package ModelTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
)

type EncoderPinModelMock struct{}

func NewEncoderPinMock(id string) Pin.EncoderPin {
	return &EncoderPinModelMock{}
}

func (encoderPin *EncoderPinModelMock) Read() bool {
	return true
}

func (encoderPin *EncoderPinModelMock) TearDown() {}

func (encoderPin *EncoderPinModelMock) WaitForEdge() {}
