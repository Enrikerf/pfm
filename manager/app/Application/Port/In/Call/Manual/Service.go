package Manual

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/CallPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
)

type Service struct {
	CallBidiPort   CallPort.Bidirectional
	FindTaskPort   TaskPort.Find
	SaveResultPort ResultPort.Save
}

func New(
	callBidiPort CallPort.Bidirectional,
	findTaskPort TaskPort.Find,
	saveResultPort ResultPort.Save,
) *Service {
	return &Service{
		CallBidiPort:   callBidiPort,
		FindTaskPort:   findTaskPort,
		SaveResultPort: saveResultPort,
	}
}

func (service *Service) ExecuteTask(task *Entity.Task, batchUuid uuid.UUID) {
	if len(task.Steps) < 0 {
		return
	}
	err := service.CallBidiPort.Setup(task.Host, task.Port)
	if err != nil {
		return
	}
	if !service.write(task.Steps[0]) {
		service.CallBidiPort.Close()
		return
	}
	for {
		taskNow, _ := service.FindTaskPort.Find(task.Uuid.String())
		if taskNow.Status != ValueObject.Running {
			service.write(task.Steps[1])
			break
		}
		resultsContent, err := service.CallBidiPort.Read()
		if err != nil {
			service.write(task.Steps[1])
			break
		}
		if (ValueObject.ResultVo{}) != resultsContent {
			result := Entity.FromResultVo(batchUuid, resultsContent)
			err := service.SaveResultPort.Save(result)
			if err != nil {
				service.write(task.Steps[1])
				break
			}
		}
	}
	service.CallBidiPort.Close()
}

func (service *Service) write(step Entity.Step) bool {
	err := service.CallBidiPort.Write(step)
	if err != nil {
		return false
	}
	return true
}
