package Entity

import (
	"errors"
	"github.com/Enrikerf/pfm/commandExecutor/app/Domain/Entity/Converter"
)

type RadiansPerSecond float64
type PWMDuty float64

type ControlAlgorithm interface {
	SetGoal(goal RadiansPerSecond)
	SetP(p float64)
	SetI(i float64)
	SetD(d float64)
	SetInMin(inMin RadiansPerSecond)
	SetInMax(inMax RadiansPerSecond)
	SetOutMin(outMin PWMDuty)
	SetOutMax(outMax PWMDuty)
	SetSampleTime(d float64)
	Calculate(currentValue RadiansPerSecond) (PWMDuty, error)
}

type controlAlgorithm struct {
	goal         RadiansPerSecond
	P            float64
	I            float64
	D            float64
	inMin        RadiansPerSecond
	inMax        RadiansPerSecond
	outMin       PWMDuty
	outMax       PWMDuty
	integralTerm float64
	sampleTime   float64
	currentValue RadiansPerSecond
	currentError RadiansPerSecond
	pastError    RadiansPerSecond
}

func (ca *controlAlgorithm) Reset() {
	ca.integralTerm = 0
	ca.currentValue = 0
	ca.currentError = 0
	ca.pastError = 0
}

func (ca *controlAlgorithm) Calculate(currentValue RadiansPerSecond) (PWMDuty, error) {
	// log.Print(" ")
	// log.Printf("currentValue: %f", ca.currentValue)
	// log.Printf("currentError: %f", ca.currentError)
	// log.Printf("goal: %f", ca.goal)
	// log.Printf("inMin: %f", ca.getMinPid())
	// log.Printf("inMax: %f", ca.getMaxPid())
	// log.Printf("outMin: %f", ca.outMin)
	// log.Printf("outMax: %f", ca.outMax)

	ca.currentValue = currentValue
	ca.currentError = ca.goal - ca.currentValue
	proportionalTerm := ca.P * float64(ca.currentError)
	ca.integralTerm = ca.integralTerm + float64(ca.currentError)*ca.sampleTime
	derivativeTerm := (float64(ca.currentError) - float64(ca.pastError)) / ca.sampleTime
	ca.pastError = ca.currentError
	pid := proportionalTerm + ca.I*ca.integralTerm + ca.D*derivativeTerm

	// log.Printf("propTerm: %f, integral: %f, derivative: %f", proportionalTerm, ca.integralTerm, derivativeTerm)
	// log.Printf("pid: %f", pid)

	if pid < ca.getMinPid() { // trash value on transitory
		return ca.outMin, nil
	}
	if pid > ca.getMaxPid() { // trash value on transitory
		return ca.outMax, nil
	}
	pidR, err := Converter.MapBetweenRanges(pid, ca.getMinPid(), ca.getMaxPid(), float64(ca.outMin), float64(ca.outMax))
	if err != nil {
		return 0, errors.New("pid out of range")
	}
	// log.Printf("pidR: %f", pidR)

	return PWMDuty(pidR), nil
}

func (ca *controlAlgorithm) getMinPid() float64 {
	return ca.P*float64(ca.inMin) + ca.I*float64(ca.inMin)*1 + ca.D*float64(ca.inMin)
}
func (ca *controlAlgorithm) getMaxPid() float64 {
	return ca.P*float64(ca.inMax) + ca.I*float64(ca.inMax)*1 + ca.D*float64(ca.inMax)
}

func NewControlAlgorithm() ControlAlgorithm {
	return &controlAlgorithm{}
}

func (ca *controlAlgorithm) SetGoal(goal RadiansPerSecond) {
	ca.goal = goal
}

func (ca *controlAlgorithm) SetP(p float64) {
	ca.P = p
}

func (ca *controlAlgorithm) SetI(i float64) {
	ca.I = i
}

func (ca *controlAlgorithm) SetD(d float64) {
	ca.D = d
}

func (ca *controlAlgorithm) SetOutMin(outMin PWMDuty) {
	ca.outMin = outMin
}

func (ca *controlAlgorithm) SetOutMax(outMax PWMDuty) {
	ca.outMax = outMax
}

func (ca *controlAlgorithm) SetInMin(inMin RadiansPerSecond) {
	ca.inMin = inMin
}
func (ca *controlAlgorithm) SetInMax(inMax RadiansPerSecond) {
	ca.inMax = inMax
}

func (ca *controlAlgorithm) SetSampleTime(st float64) {
	ca.sampleTime = st
}
