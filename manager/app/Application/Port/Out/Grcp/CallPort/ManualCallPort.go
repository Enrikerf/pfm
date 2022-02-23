package CallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Bidirectional interface {
	Setup(host, port string) error
	Write(step Entity.Step) error
	Read() (ValueObject.ResultVo, error)
	Close()
}
