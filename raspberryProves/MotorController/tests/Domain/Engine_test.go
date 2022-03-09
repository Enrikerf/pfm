package DomainTest

import (
	"testing"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
	ModelTest "github.com/Enrikerf/pfm/commands/MotorController/tests/Adapter/Out/Periphio/Model"
)

var tests = []struct{ 
	value float64 
	result int32 
	}{
	{
		
	}
}
func TestReescalatePid(t *testing.T) {

	pwmPin := ModelTest.NewPwmPinMock()
	brakePin := ModelTest.NewOutPinMock()
	dirPin := ModelTest.NewOutPinMock()
	encoder := ModelTest.NewEncoderMock()
	control := Entity.NewControlAlgorithm()
	engine := Entity.NewEngine(
		360,
		control,
		brakePin,
		dirPin,
		pwmPin,
		encoder,
	)
	engine.ReescalePid(12)
}
