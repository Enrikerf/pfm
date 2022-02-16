package main

import "fmt"

func main() {
	embed := NewSingleton("hello")
	s1 := NewUseSingleton(embed)
	s2 := NewUseSingleton(embed)
	s1.ChangeToSay("no Hello")
	fmt.Println(s1.Say())
	fmt.Println(s2.Say())

	embed1 := NewSingleton("hello")   //will no change because return instance which it's previous changed to no hello
	embed2 := NewSingleton("noHello") //will no change because return instance which it's previous changed to no hello
	s3 := NewUseSingleton(embed1)
	s4 := NewUseSingleton(embed2)
	fmt.Println(s3.Say())
	fmt.Println(s4.Say())
}
