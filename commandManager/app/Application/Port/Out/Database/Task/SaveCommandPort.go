package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type SaveCommandPort interface {
	Save(command Task.Command) error
}
