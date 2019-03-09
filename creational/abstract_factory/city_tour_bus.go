package abstract_factory

// A group of Bus family
type CityTourBus struct{
	wheels int
	doors int
	speed int
	floorCount int
	
	// Some CityTourBus specific properties
}

func (c CityTourBus) FloorCount() int {
	return c.floorCount
}

func (c CityTourBus) WheelCount() int {
	return c.wheels
}

func (c CityTourBus) NumberOfDoors() int {
	return c.doors
}

func (c CityTourBus) Speed() int {
	return c.speed
}