package abstract_factory

import (
	"fmt"
	"errors"
)

const (
	MICRO_BUS = 1
	CITY_TOUR_BUS = 2
)

type BusFactory struct {}

// Creates vehicles which is member of Bus subfamily group
func(f *BusFactory) NewVehicle(busType int)(Vehiche, error) {
	switch busType {
	case CITY_TOUR_BUS:
		return new(CityTourBus), nil
	case MICRO_BUS:
		return new(MicroBus), nil
	default: 
		return nil, errors.New(fmt.Sprintf("Unsupported bus vehicle type:%d", busType))
	}
}
