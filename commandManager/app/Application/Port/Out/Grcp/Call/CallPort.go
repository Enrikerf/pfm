package Call

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type RequestPort interface {
	Request(result Result.Result) error
}
