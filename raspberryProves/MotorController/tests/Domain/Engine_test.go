package DomainTest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
	ModelTest "github.com/Enrikerf/pfm/commands/MotorController/tests/Adapter/Out/Periphio/Model"
)

/**
conditions
pid (-inf,inf)
min (0,inf)
max [0,inf)
min>max

TEST CASES
case    PID     MIN     rel				result
1       pos     0       pid<min<max		no sense positive number greater than 0
2       pos     pos     pid>min<max		pid
3       neg     pos     pid<min<max		min
4       neg     0       pid>min<max		no sense negative number greater than 0
5       0       0       pid>min>max		no sense max must be greater than min and min=pid=0
6       0       pos     pid<min<max		min
7       pos     pos     pid>min>max		max
8       neg     ~0      pid>min>max		no sense negative number greater than 0
9       0       ~0      pid>min<max		no sense
*/

const pidNeg = -10
const pidPos1 = 10
const pidPos2 = 12
const zero = 0
const minPos = 9
const max = 11

/*
Tests to implement
case    PID     MIN     rel				result
1       pos     pos     pid>min<max		pid
2       neg     pos     pid<min<max		min
3       0       pos     pid<min<max		min
4       pos     pos     pid>min>max		max
*/
var tests = []struct {
	pid    float64
	min    int32
	max    int32
	result int32
}{
	{
		pid:    pidPos1,
		min:    minPos,
		max:    max,
		result: pidPos1,
	},
	{
		pid:    pidNeg,
		min:    minPos,
		max:    max,
		result: minPos,
	},
	{
		pid:    zero,
		min:    minPos,
		max:    max,
		result: minPos,
	},
	{
		pid:    pidPos2,
		min:    minPos,
		max:    max,
		result: max,
	},
}

func TestReescalatePid(t *testing.T) {
	for _, test := range tests {
		pwmPin := ModelTest.NewPwmPinMock(test.max, test.min)
		brakePin := ModelTest.NewOutPinMock()
		dirPin := ModelTest.NewOutPinMock()
		encoder := ModelTest.NewEncoderMock()
		control := Entity.NewControlAlgorithm()
		engine := Entity.NewEngine(
			1.5,
			200,
			360,
			control,
			brakePin,
			dirPin,
			pwmPin,
			encoder,
		)
		assert.Equal(t, test.result, engine.ConstrainInDuty(test.pid), "must be equals")
	}
}
