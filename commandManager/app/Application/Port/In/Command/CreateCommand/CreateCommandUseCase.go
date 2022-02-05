package CreateCommand

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UseCase interface {
	Create(command Command) (TaskDomain.Command, error)
}
