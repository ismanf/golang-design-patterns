package abstract_factory

// A group of Car family
type Sportcar struct{
	wheels int
	doors int
	speed int
	hasElectricEngine bool

	// Some sportcar specific properties
}

func (c Sportcar) WheelCount() int {
	return c.wheels
}

func (c Sportcar) NumberOfDoors() int {
	return c.doors
}

func (c Sportcar) Speed() int {
	return c.speed
}

func (c Sportcar) HasElectricEngine() bool {
	return c.hasElectricEngine
}