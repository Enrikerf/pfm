package ManageEngine

import "github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"

type Service struct {
	Engine Entity.Engine
}

func (service *Service) Turnaround() {
	service.Engine.MakeLap()
}

func (service *Service) RpmControl(rpm float64) {
	service.Engine.RpmControl(rpm)
}

func (service *Service) StopRpmControl() {
	service.Engine.StopRmpControl()
}

func (service *Service) SetGas(gas int) {
	service.Engine.SetGas(Entity.GasLevel(gas))
}

func (service *Service) StepResponse() {
	service.Engine.StepResponse()
}

func (service *Service) UnBrake() {
	service.Engine.UnBrake()
}

func (service *Service) TearDown() {
	service.Engine.TearDown()
}
