package Repository

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"

type Delete interface {
	Delete(id Task.Id) error
}
