package abstract_factory

/**
*  Car is a subfamliy of Vehicle 
*  But it has some specific properties 
*  thus based on these properties has some variations
*  Therefore it need it's own factory
*/
type Car interface {
	
	// Just embeding Vehicle interface 
	// in order to extend from it
	Vehiche

	// Let's assume this property belongs to
	// only Car type vehicles
	HasElectricEngine() bool
}
