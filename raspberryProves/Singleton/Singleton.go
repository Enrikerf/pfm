package main

import "sync"

var once sync.Once

var instance *singleton

type Singleton interface {
	ChangeToSay(toSay string)
	Say() string
}
type singleton struct {
	toSay string
}

func (singleton *singleton) ChangeToSay(toSay string) {
	singleton.toSay = toSay
}

func (singleton *singleton) Say() string {
	return singleton.toSay
}

func NewSingleton(toSay string) Singleton {

	once.Do(func() {
		instance = &singleton{
			toSay: toSay,
		}
	})
	return instance
}
