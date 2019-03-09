package abstract_factory

// A group of Car family
type Offroader struct{
	wheels int
	doors int
	speed int
	hasElectricEngine bool
	
	// Some offroader specific properties
}

func (c Offroader) WheelCount() int {
	return c.wheels
}

func (c Offroader) NumberOfDoors() int {
	return c.doors
}

func (c Offroader) Speed() int {
	return c.speed
}

func (c Offroader) HasElectricEngine() bool {
	return c.hasElectricEngine
}