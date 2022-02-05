package ListCommands

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UseCase interface {
	List(query Query) []TaskDomain.Command
}
