package main

type UseSingleton interface {
	ChangeToSay(toSay string)
	Say() string
}
type useSingleton struct {
	singleton Singleton
}

func (useSingleton *useSingleton) ChangeToSay(toSay string) {
	useSingleton.singleton.ChangeToSay(toSay)
}

func (useSingleton *useSingleton) Say() string {
	return useSingleton.singleton.Say()
}

func NewUseSingleton(singleton Singleton) UseSingleton {
	return &useSingleton{singleton: singleton}
}
