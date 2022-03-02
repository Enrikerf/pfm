package ListResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindResultsByPort ResultPort.FindBy
}

func (service Service) List(query Query) []Entity.Result {
	conditions := map[string]interface{}{}
	if query.BatchUuid.Change {
		conditions = map[string]interface{}{"batch_uuid": query.BatchUuid.Value}
	}
	return service.FindResultsByPort.FindBy(conditions)
}
