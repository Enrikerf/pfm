package ReadCommand

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UseCase interface {
	Read(query Query) (TaskDomain.Command, error)
}
