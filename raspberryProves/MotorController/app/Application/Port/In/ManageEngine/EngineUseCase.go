package ManageEngine

type UseCase interface {
	Turnaround()
	RpmControl(rpm float64)
	StopRpmControl()
	UnBrake()
	SetGas(gas int)
	StepResponse()
	TearDown()
}
