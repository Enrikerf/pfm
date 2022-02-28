package UpdateResult

import "github.com/Enrikerf/pfm/commandManager/app/Application/Model"

type Command struct {
	Uuid      string
	Content   Model.OptionalString
	BatchUuid Model.OptionalString
}
