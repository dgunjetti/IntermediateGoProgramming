package main

import "fmt"

// START1 OMIT
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() Vehicle
}

// END1 OMIT

// START2 OMIT
type CarBuilder struct {
	v Vehicle
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 3
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() Vehicle {
	return c.v
}

// END2 OMIT

//START3 OMIT
type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}
func (d *Director) Construct() {
	d.builder.SetSeats().SetWheels().SetStructure()
}
func main() {
	d := Director{}
	cb := &CarBuilder{}
	d.SetBuilder(cb)
	d.Construct()
	c := cb.GetVehicle()
	fmt.Println(c)
}

//END3 OMIT
