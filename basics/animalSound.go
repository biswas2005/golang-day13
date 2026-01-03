package basics

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}
type Cat struct{}

func (s Dog) Sound() string {
	return "Bhau"
}
func (s Cat) Sound() string {
	return "Meow"
}

func AnimalSound() {
	d := Dog{}
	c := Cat{}

	sound := []Animal{d, c}

	for _, val := range sound {
		fmt.Printf("Sound of %T is %s\n", val, val.Sound())
	}
}
