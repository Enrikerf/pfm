package ListResults

import (
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindResultsByPort ResultOutPort.FindResultsByPort
}

func (service Service) List(query Query) []Result.Result {
	return service.FindResultsByPort.FindBy(query)
}
