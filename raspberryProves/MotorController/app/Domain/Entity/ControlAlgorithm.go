package Entity

import (
	"log"
	"math"
)

type ControlAlgorithm interface {
	SetGoal(goal float64)
	SetP(p float64)
	SetI(i float64)
	SetD(d float64)
	SetLowerConstraint(constraint float64)
	SetUpperConstraint(constraint float64)
	SetSampleTime(d float64)
	GetIntegralTerm() float64
	Calculate(currentValue float64) float64
}

type controlAlgorithm struct {
	goal            float64
	P               float64
	I               float64
	D               float64
	lowerConstraint float64
	upperConstraint float64
	integralTerm    float64
	sampleTime      float64
	currentValue    float64
	currentError    float64
	pastError       float64
}

func (ca *controlAlgorithm) Calculate(currentValue float64) float64 {
	ca.currentValue = currentValue

	x := ca.rpm2radps(1.5)
	y := ca.rpm2radps(200)
	currentValueMapped := ca.mapBetweenRanges(ca.currentValue, x, y, ca.lowerConstraint, ca.upperConstraint)
	goalMapped := ca.mapBetweenRanges(ca.goal, x, y, ca.lowerConstraint, ca.upperConstraint)
	ca.currentError = goalMapped - currentValueMapped
	// ca.currentError = ca.mapBetweenRanges(ca.goal-ca.currentValue, x, y, ca.lowerConstraint, ca.upperConstraint)
	proportionalTerm := ca.P * ca.currentError
	ca.integralTerm = ca.integralTerm + ca.currentError*ca.sampleTime
	derivativeTerm := ca.D * (ca.currentError - ca.pastError) / ca.sampleTime
	ca.pastError = ca.currentError
	pid := proportionalTerm + ca.I*ca.integralTerm + derivativeTerm
	log.Print(" ")
	log.Printf("currentValue: %f", ca.currentValue)
	log.Printf("goal: %f", ca.goal)
	log.Printf("x: %f", x)
	log.Printf("y: %f", y)
	log.Printf("low: %f", ca.lowerConstraint)
	log.Printf("up: %f", ca.upperConstraint)
	log.Printf("currentValueMapped: %f", currentValueMapped)
	log.Printf("goalMapped: %f", goalMapped)
	log.Printf("currentError: %f", ca.currentError)
	log.Printf("propTerm: %f, integral: %f, derivative: %f", proportionalTerm, ca.integralTerm, derivativeTerm)
	log.Printf("pid: %f", pid)
	return goalMapped
}

func (ca *controlAlgorithm) radps2rpm(radps float64) float64 {
	return radps * 60 / (2 * math.Pi)
}

func (ca *controlAlgorithm) rpm2radps(rpm float64) float64 {
	return rpm * (2 * math.Pi) / 60
}

func (ca *controlAlgorithm) degps2Radps(degreesPerSecod float64) float64 {
	return degreesPerSecod * math.Pi / 180
}

func (ca *controlAlgorithm) mapBetweenRanges(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x - inMin) * (outMax - outMin) / (inMax - inMin) + outMin; 
}

func (ca *controlAlgorithm) Reescale(value float64) float64 {
	if value < float64(ca.lowerConstraint) {
		return ca.lowerConstraint
	}
	if value > ca.upperConstraint {
		return ca.upperConstraint
	}
	return value
}

func NewControlAlgorithm() ControlAlgorithm {
	return &controlAlgorithm{}
}

func (ca *controlAlgorithm) SetGoal(goal float64) {
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
func (ca *controlAlgorithm) SetLowerConstraint(constraint float64) {
	ca.lowerConstraint = constraint
}
func (ca *controlAlgorithm) SetUpperConstraint(constraint float64) {
	ca.upperConstraint = constraint
}

func (ca *controlAlgorithm) GetIntegralTerm() float64 {
	return ca.integralTerm
}

func (ca *controlAlgorithm) SetSampleTime(st float64) {
	ca.sampleTime = st
}
