package Manual

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/CallPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
	"time"
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
	service.CallBidiPort.Setup()
	service.CallBidiPort.Write(task.Steps[0])
	for {
		taskNow, _ := service.FindTaskPort.Find(task.Uuid.String())
		if taskNow.Status != ValueObject.Running {
			service.CallBidiPort.Write(task.Steps[1])
			break
		}
		resultsContent := service.CallBidiPort.Read()
		if (ValueObject.ResultVo{}) != resultsContent {
			result := Entity.FromResultVo(batchUuid, resultsContent)
			err := service.SaveResultPort.Save(result)
			if err != nil {
				service.CallBidiPort.Write(task.Steps[1])
				break
			}
		}
		time.Sleep(time.Second)
	}
}
