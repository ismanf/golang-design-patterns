package abstract_factory

/**
*  Vehicle is root abstractions for the family 
*  It contains very common behavior for the
*  family tree
*/
type Vehiche interface {

	WheelCount() int
	
	NumberOfDoors() int
	
	Speed() int
}