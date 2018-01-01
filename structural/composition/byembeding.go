package composition

import (
	"fmt"
)

type Car struct {
	WheelCount uint8
	Speed int
	DoorCount uint8
	Brand string
}

func (c *Car) Drive() {
	fmt.Println("Goes with speed: %d", c.Speed)
}

type FlyingCar struct {
	Car
	WingLength float32
	FlySpeed int
} 

func (c *FlyingCar) Fly()  {
	fmt.Println("Flies with speed: %d", c.FlySpeed)
}

type SwimmingCar struct {
	Car
	SwimSpeed int
}

func (c *SwimmingCar) Swimm()  {
	fmt.Println("Swimm with speed: %d", c.SwimSpeed)
}