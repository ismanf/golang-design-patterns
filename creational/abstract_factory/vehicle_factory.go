package abstract_factory

/**
*  Is abstract factory for Vehicle family groups
*  Each new family group factory should implement
*  this interface
*/
type VehicleFactory interface {
	NewVehicle(int) (Vehiche, error)
}