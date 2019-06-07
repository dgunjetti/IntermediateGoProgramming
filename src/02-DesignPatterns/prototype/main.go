package main

import (
	"fmt"
	"log"
)

//START1 OMIT
const (
	White = 1
	Black = 2
)

type ShirtsCloner interface {
	GetClone(s int) (ItemsInfoGetter, error)
}
type ShirtsCache struct{}

func (*ShirtsCache) GetClone(s int) (ItemsInfoGetter, error) {
	switch s {
	case White: // can add cases for black, blue
		newItem := whitePrototype
		return newItem, nil
	default:
		return nil, fmt.Errorf("Invalid: %d", s)
	}
}

//END1 OMIT

//START2 OMIT
type ItemsInfoGetter interface {
	GetInfo() string
}
type Shirt struct {
	Price float64
	Color int
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("%d, %0.2f", s.Color, s.Price)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	Color: White,
}

// can add blackPrototype

//END2 OMIT
//START3 OMIT
func main() {
	sc := &ShirtsCache{}
	w, err := sc.GetClone(White)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w.GetInfo())
}

//END3 OMIT
