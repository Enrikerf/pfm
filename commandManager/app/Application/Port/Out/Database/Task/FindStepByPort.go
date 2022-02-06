package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindStepByPort interface {
	FindBy(conditions interface{}) []Task.Step
}
