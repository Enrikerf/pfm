package TaskPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindBy interface {
	FindBy(conditions interface{}) []Task.Task
}
