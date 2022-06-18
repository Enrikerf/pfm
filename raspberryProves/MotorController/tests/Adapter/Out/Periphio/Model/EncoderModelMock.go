package ModelTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
)

type EncoderModelMock struct{}

func NewEncoderMock() Entity.Encoder {
	return &EncoderModelMock{}
}

func (encoder *EncoderModelMock) GetPosition() int16 { return 0 }

func (encoder *EncoderModelMock) ResetPosition() {}

func (encoder *EncoderModelMock) TearDown() {}

func (encoder *EncoderModelMock) Watchdog() {}
