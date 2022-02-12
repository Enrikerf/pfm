package Entity

type ControlAlgorithm interface {
	SetGoal(goal float64)
	SetP(p float64)
	SetI(i float64)
	SetD(d float64)
	Calculate(currentValue float64) float64
}

type controlAlgorithm struct {
	goal         float64
	P            float64
	I            float64
	D            float64
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

func (ca *controlAlgorithm) Calculate(currentValue float64) float64 {
	ca.currentValue = currentValue
	ca.currentError = ca.goal - ca.currentValue
	integralTerm := ca.I * ca.currentError
	derivativeTerm := ca.D * (ca.currentError - ca.pastError)
	proportionalTerm := ca.P * ca.currentError
	ca.pastError = ca.currentError
	return proportionalTerm + integralTerm + derivativeTerm

}
