package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type FindByPort interface {
	FindBy(conditions interface{}) []Result.Result
}
