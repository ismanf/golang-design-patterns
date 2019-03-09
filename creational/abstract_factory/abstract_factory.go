package abstract_factory

import (
	"fmt"
	"errors"
)

const (
	CAR = 1
	BUS = 2
)

/**
*  CreateVehicleFactory function makes subfamily factories of Vehicle
*   family. The goal is to split object family creation complexity
*   into small blocks in order to make it easy to read and maintain.
*   It basicaly delegates subfamily object creations to subfamily factories.
*/
func CreateVehicleFactory(vfType int) (VehicleFactory, error) {
	switch vfType {
	case CAR:
		return new(CarFactory), nil
	case BUS:
		return new(BusFactory), nil
	default: 
		return nil, errors.New(fmt.Sprintf("Unrecognized factory type:%d", vfType))
	}
}