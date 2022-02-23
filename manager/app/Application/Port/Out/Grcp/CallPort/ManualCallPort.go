package CallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Bidirectional interface {
	Setup()
	Write(step Entity.Step)
	Read() ValueObject.ResultVo
}
