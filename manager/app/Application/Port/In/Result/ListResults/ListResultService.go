package ListResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindResultsByPort ResultPort.FindBy
}

func (service Service) List(query Query) []Entity.Result {
	return service.FindResultsByPort.FindBy(query)
}
