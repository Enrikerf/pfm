package BatchPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type FindBy interface {
	FindBy(conditions interface{}) []Result.Batch
}
