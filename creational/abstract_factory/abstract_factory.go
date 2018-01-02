package abstract_factory

import (
	"fmt"
	"errors"
)

//sub factory abstractions
type VehicleFactory interface {
	Create(int) (Vehiche, error)
}

type CarFactory struct{}
func (f *CarFactory) Create(vType int) (Vehiche, error) {
	switch vType {
	case OFFROADER_CAR:
		return new(Offroader), nil
	case SPORT_CAR:
		return new(Sportcar), nill
	default: 
		return nil, errors.New(fmt.Sprintf("Unsupported car vehicle type:%d", vType))
	}
}

type BusFactory struct {}
func(f *BusFactory) Create(vType int)(Vehiche, error) {
	switch vType {
	case CITY_TOUR_BUS:
		return new(CityTourBus), nil
	case SIMPLE_BUS:
		return new(SimpleBus), nil
	default: 
		return nil, errors.New(fmt.Sprintf("Unsupported bus vehicle type:%d", vType))
	}
}

//object families............
type Vehiche interface {
	WheelCount() int
}

//Car
const (
	OFFROADER_CAR = 1
	SPORT_CAR = 2
)

type Car interface {
	DoorCount() int
}

type Offroader struct{}
func (c Offroader) WheelCount() int {
	return 4
}
func (c Offroader) DoorCount() int {
	return 4
}

type Sportcar struct{}
func (c Sportcar) WheelCount() int {
	return 4
}
func (c Sportcar) DoorCount() int {
	return 2
}

//Bus..................
const (
	SIMPLE_BUS = 1
	CITY_TOUR_BUS = 2
)

type Bus interface {
	FloorCount() int
}

type CityTourBus struct{}
func (c CityTourBus) WheelCount() int {
	return 6
}
func (c CityTourBus) FloorCount() int {
	return 2
}

type SimpleBus struct{}
func (c SimpleBus) WheelCount() int {
	return 4
}
func (c SimpleBus) FloorCount() int {
	return 1
}

//super factory
const (
	CAR = 1
	BUS = 2
)

func CreateVehicleFactory(vfType int) (VehicleFactory, error) {
	switch vfType {
	case CAR:
		return new(CarFactory), nil
	case BUS:
		return new(BusFactory), nil
	default: 
		return nil, errors.New("Unrecognized factory type:%d", vfType)
	}
}