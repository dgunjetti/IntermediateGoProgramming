package main

//START1 OMIT
type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Swimmer struct{}

func (s *Swimmer) Swim() {
	println("Swimming")
}

type Shark struct {
	A Animal
	S Swimmer
}

//END1 OMIT

//START2 OMIT
func main() {
	s := &Shark{}
	s.A.Eat()
	s.S.Swim()
}

//END2 OMIT
