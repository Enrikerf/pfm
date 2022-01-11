package LoopManager

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Loop"
)

type LoopManager struct {
	Loop Loop.UseCase
}

func (loopManager LoopManager) Execute() {
	err := loopManager.Loop.Loop()
	if err != nil {
		fmt.Println(err)
	}
}
