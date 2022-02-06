package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type FindResultsByPort interface {
	FindBy(conditions interface{}) []Result.Result
}
