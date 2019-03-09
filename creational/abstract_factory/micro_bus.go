package abstract_factory

// A group of Bus family
type MicroBus struct{
	wheels int
	doors int
	speed int
	floorCount int
	
	// Some CityTourBus specific properties
}

func (c MicroBus) FloorCount() int {
	return c.floorCount
}

func (c MicroBus) WheelCount() int {
	return c.wheels
}

func (c MicroBus) NumberOfDoors() int {
	return c.doors
}

func (c MicroBus) Speed() int {
	return c.speed
}