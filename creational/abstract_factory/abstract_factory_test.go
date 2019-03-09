package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T)  {
	
	carFactory, err := CreateVehicleFactory(CAR)
	if err != nil {
		t.Fatal(err)
	}

	sportCarVehicle, err := carFactory.NewVehicle(SPORT_CAR)
	if err != nil {
		t.Fatal(err)
	}

	sc, ok := sportCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}

func TestBusFactory(t *testing.T)  {
	
	busFactory, err := CreateVehicleFactory(BUS)
	if err != nil {
		t.Fatal(err)
	}

	cistyTourBusVehicle, err := busFactory.NewVehicle(CITY_TOUR_BUS)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := cistyTourBusVehicle.(Bus)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}