package ListResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindResultsByPort ResultPort.FindBy
}

func (service Service) List(query Query) []Result.Result {
	return service.FindResultsByPort.FindBy(query)
}
