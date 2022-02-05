package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindCommandByPort interface {
	FindBy(conditions interface{}) []Task.Command
}
