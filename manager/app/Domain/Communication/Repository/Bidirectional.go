package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Bidirectional interface {
	Setup(host, port string) error
	Write(step Step.Step) error
	Read() (Content.Content, error)
	Close()
}
