package StepPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type Find interface {
	Find(uuid string) (Task.Step, error)
}
