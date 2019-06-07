package main

import "fmt"

type Trainer struct{}
type Animal struct{}

//START1 OMIT
type Athlete struct {
	Trainer
	Swim func()
}

type Fish struct {
	Animal
	Swim func()
}

func swimfunc() {
	fmt.Println("swim")
}

//END1 OMIT

//START2 OMIT
func main() {
	a := Athlete{
		Swim: swimfunc,
	}
	a.Swim()

	f := Fish{
		Swim: swimfunc,
	}
	f.Swim()
}

//END2 OMIT
