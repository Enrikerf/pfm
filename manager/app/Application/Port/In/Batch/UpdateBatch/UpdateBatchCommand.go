package UpdateBatch

import "github.com/Enrikerf/pfm/commandManager/app/Application/Model"

type Command struct {
	Uuid     string
	TaskUuid Model.OptionalString
}
