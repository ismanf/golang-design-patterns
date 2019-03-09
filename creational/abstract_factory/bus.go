package abstract_factory

/**
*  Bus is also a subfamliy of Vehicle 
*  But it has some specific properties 
*  thus based on these properties has some variations
*  Therefore it need it's own factory
*/
type Bus interface {

    Vehiche
	
	// As some buses can have multiple floors
	// let's consider this is a big difference from
	// other vehicles
	FloorCount() int
}