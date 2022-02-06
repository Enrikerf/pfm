package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID            uint
	Uuid          uuid.UUID
	Host          string
	Port          string
	Steps         []Step
	Mode          string
	Status        string
	ExecutionMode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Task) TableName() string {
	return "tasks"
}

func (taskModel *Task) FromDomain(selfEntity Entity.Task) {
	taskModel.Uuid = selfEntity.Uuid
	taskModel.Host = selfEntity.Host
	taskModel.Port = selfEntity.Port
	for _, stepDomain := range selfEntity.Steps {
		step := Step{}
		step.FromDomain(stepDomain)
		taskModel.Steps = append(taskModel.Steps, step)
	}
	taskModel.Mode = selfEntity.Mode.String()
	taskModel.Status = selfEntity.Status.String()
	taskModel.ExecutionMode = selfEntity.ExecutionMode.String()
}

func (taskModel *Task) ToDomain() Entity.Task {
	selfEntity := Entity.Task{}
	selfEntity.Uuid = taskModel.Uuid
	selfEntity.Host = taskModel.Host
	selfEntity.Port = taskModel.Port
	for _, stepMysql := range taskModel.Steps {
		selfEntity.Steps = append(selfEntity.Steps, stepMysql.ToDomain())
	}
	selfEntity.Mode, _ = ValueObject.GetTaskMode(taskModel.Mode)
	selfEntity.Status, _ = ValueObject.GetStatus(taskModel.Status)
	selfEntity.ExecutionMode, _ = ValueObject.GetExecutionMode(taskModel.ExecutionMode)
	return selfEntity
}
