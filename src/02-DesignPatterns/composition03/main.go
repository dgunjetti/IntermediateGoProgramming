package main

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Swimmer struct{}

func (s *Swimmer) Swim() {
	println("Swimming")
}

//START1 OMIT
type Shark struct {
	Animal
	Swimmer
}

//END1 OMIT

//START2 OMIT
func main() {
	s := &Shark{}
	s.Eat()
	s.Swim()
}

//END2 OMIT
