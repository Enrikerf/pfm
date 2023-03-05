package Entity
import	"fmt"

type ControlAlgorithm interface {
	SetGoal(goal float64)
	SetP(p float64)
	SetI(i float64)
	SetD(d float64)
	SetSampleTime(d float64)
	GetSampleTime() float64
	GetIntegralTerm() float64
	Calculate(currentValue float64) float64
}

type controlAlgorithm struct {
	goal         float64
	P            float64
	I            float64
	D            float64
	integralTerm float64
	sampleTime   float64
	currentValue float64
	currentError float64
	pastError    float64
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

func (ca *controlAlgorithm) GetIntegralTerm() float64 {
	return ca.integralTerm
}

func (ca *controlAlgorithm) SetSampleTime(st float64) {
	ca.sampleTime = st
}

func (ca *controlAlgorithm) GetSampleTime()float64 {
	return ca.sampleTime 
}

func (ca *controlAlgorithm) Calculate(currentValue float64) float64 {
	fmt.Println("\tsubfuc")
	fmt.Println("\t\tca.goalg:",ca.goal)
	ca.currentValue = currentValue
	fmt.Println("\t\tcurrentValue:",currentValue)
	ca.currentError = ca.goal - ca.currentValue
	fmt.Println("\t\tca.currentError:",ca.currentError)
	proportionalTerm := ca.P * ca.currentError
	fmt.Println("\t\tproportionalTerm:",proportionalTerm)
	ca.integralTerm = ca.integralTerm + ca.I*ca.currentError*ca.sampleTime
	fmt.Println("\t\tintegralTerm:",ca.integralTerm)
	derivativeTerm := ca.D * (ca.currentError - ca.pastError) / ca.sampleTime
	// fmt.Println("\t\tderivativeTerm:",derivativeTerm)
	ca.pastError = ca.currentError
	total :=  proportionalTerm + ca.integralTerm + derivativeTerm
	fmt.Println("\t\ttotal:",total)
	return total

}
