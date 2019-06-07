package main

import (
	"fmt"
	"log"
)

//START1 OMIT
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}
type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

//END1 OMIT
//START2 OMIT
const (
	CarFactoryType  = 1
	BikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case BikeFactoryType:
		return new(BikeFactory), nil
	default:
		return nil, fmt.Errorf("%d not registered", f)
	}
}

//END2 OMIT
//START3 OMIT
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type Car interface {
	NumDoors() int
}
type CarFactory struct{}

func (cf *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("%d not registered", v)
	}
}

//END3 OMIT
//START4 OMIT
type LuxuryCar struct{}

func (l *LuxuryCar) NumWheels() int {
	return 4
}
func (l *LuxuryCar) NumSeats() int {
	return 5
}
func (l *LuxuryCar) NumDoors() int {
	return 4
}

//END4 OMIT

type FamilyCar struct{}

func (l *FamilyCar) NumWheels() int {
	return 5
}
func (l *FamilyCar) NumSeats() int {
	return 5
}
func (l *FamilyCar) NumDoors() int {
	return 4
}

//START5 OMIT
const (
	SportsType = 1
	CruiseType = 2
)

type Bike interface {
	GetBikeType() int
}
type BikeFactory struct{}

func (bf *BikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportsType:
		return new(Sports), nil
	case CruiseType:
		return new(Cruise), nil
	default:
		return nil, fmt.Errorf("%d not registered\n", v)
	}
}

//END5 OMIT

type Sports struct{}

func (l *Sports) NumWheels() int {
	return 2
}
func (l *Sports) NumSeats() int {
	return 1
}
func (l *Sports) GetBikeType() string {
	return "Sports"
}

type Cruise struct{}

func (l *Cruise) NumWheels() int {
	return 2
}
func (l *Cruise) NumSeats() int {
	return 2
}
func (l *Cruise) GetBikeType() string {
	return "Cruise"
}

//START6 OMIT
func main() {
	cf, err := BuildFactory(CarFactoryType)
	if err != nil {
		log.Fatalf("Build factory return error: %v", err)
	}
	c, _ := cf.Build(LuxuryCarType)

	l, ok := c.(Car)
	if !ok {
		log.Fatal("type assertion on car failed")
	}
	fmt.Println(l.NumDoors())
}

//END6 OMIT
