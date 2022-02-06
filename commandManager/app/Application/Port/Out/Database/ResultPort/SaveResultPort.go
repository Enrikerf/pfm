package ResultPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type Save interface {
	Save(result Result.Result) error
}
