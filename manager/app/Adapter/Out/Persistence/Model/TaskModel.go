package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
	"time"
)

type TaskDb struct {
	ID            uint
	Uuid          uuid.UUID
	Host          string
	Port          string
	Steps         []StepDb `gorm:"foreignKey:TaskID;references:ID"`
	Mode          string
	Status        string
	ExecutionMode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (TaskDb) TableName() string {
	return "tasks"
}

func (taskDb *TaskDb) FromDomainV2(selfEntity TaskDomain.Task) {
	taskDb.Uuid = selfEntity.GetUuid()
	taskDb.Host = selfEntity.GetHost().GetValue()
	taskDb.Port = selfEntity.GetPort().GetValue()
	for _, stepDomain := range selfEntity.GetSteps() {
		step := StepDb{}
		step.FromDomainV2(taskDb.Uuid, stepDomain)
		taskDb.Steps = append(taskDb.Steps, step)
	}
	taskDb.Mode = string(selfEntity.GetCommunicationMode())
	taskDb.Status = string(selfEntity.GetStatus())
	taskDb.ExecutionMode = string(selfEntity.GetExecutionMode())
}

func (taskDb *TaskDb) FromDomain(selfEntity Entity.Task) {
	taskDb.Uuid = selfEntity.Uuid
	taskDb.Host = selfEntity.Host
	taskDb.Port = selfEntity.Port
	for _, stepDomain := range selfEntity.Steps {
		step := StepDb{}
		step.FromDomain(stepDomain)
		taskDb.Steps = append(taskDb.Steps, step)
	}
	taskDb.Mode = selfEntity.Mode.String()
	taskDb.Status = selfEntity.Status.String()
	taskDb.ExecutionMode = selfEntity.ExecutionMode.String()
}

func (taskDb *TaskDb) ToDomain() Entity.Task {
	selfEntity := Entity.Task{}
	selfEntity.Uuid = taskDb.Uuid
	selfEntity.Host = taskDb.Host
	selfEntity.Port = taskDb.Port
	for _, stepMysql := range taskDb.Steps {
		selfEntity.Steps = append(selfEntity.Steps, stepMysql.ToDomain())
	}
	selfEntity.Mode, _ = ValueObject.GetTaskMode(taskDb.Mode)
	selfEntity.Status, _ = ValueObject.GetStatus(taskDb.Status)
	selfEntity.ExecutionMode, _ = ValueObject.GetExecutionMode(taskDb.ExecutionMode)
	return selfEntity
}
