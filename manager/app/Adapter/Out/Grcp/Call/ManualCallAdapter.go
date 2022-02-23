package Call

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type ManualAdapter struct {
}

func (manualAdapter *ManualAdapter) Setup() {

}
func (manualAdapter *ManualAdapter) Write(step Entity.Step) {

}
func (manualAdapter *ManualAdapter) Read() ValueObject.ResultVo {
	return ValueObject.ResultVo{}
}
