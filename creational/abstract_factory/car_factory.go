package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	OFFROADER_CAR = 1
	SPORT_CAR     = 2
)

type CarFactory struct{}

// Creates vehicles which is member of Car subfamily group
func (f *CarFactory) NewVehicle(carType int) (Vehiche, error) {
	switch carType {
	case OFFROADER_CAR:
		return new(Offroader), nil
	case SPORT_CAR:
		return new(Sportcar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported car vehicle type:%d", carType))
	}
}
