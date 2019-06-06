//START1 OMIT
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//END1 OMIT

//START2 OMIT
type CarBuilder struct { 
    v VehicleProduct 
} 
 
func (c *CarBuilder) SetWheels() BuildProcess { 
    c.v.Wheels = 4 
    return c 
} 
 
func (c *CarBuilder) SetSeats() BuildProcess { 
    c.v.Seats = 5 
    return c 
} 
 
func (c *CarBuilder) SetStructure() BuildProcess { 
    c.v.Structure = "Car" 
    return c 
} 
 
func (c *CarBuilder) GetVehicle() VehicleProduct { 
    return c.v 
} 

//END2 OMIT

//START3 OMIT
package creational 
 
type ManufacturingDirector struct { 
    builder BuildProcess 
} 
 
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) { 
    f.builder = b 
} 
 
func (f *ManufacturingDirector) Construct() { 
    f.builder.SetSeats().SetStructure().SetWheels() 
} 
//END3 OMIT