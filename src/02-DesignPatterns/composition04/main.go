package main

import (
	"fmt"
)

//START1 OMIT
type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type Composite struct {
	Trainer
	Swimmer
}

//END1 OMIT

//START2 OMIT
type Athlete struct{}

func (a *Athlete) Train() { fmt.Println("train") }

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() { fmt.Println("swim") }

//END2 OMIT
func main() {
	//START3 OMIT
	c := Composite{
		Trainer: &Athlete{},
		Swimmer: &SwimmerImpl{},
	}
	//END3 OMIT
	fmt.Println(c)
}
