package builder

// Product to build
type PersonalComputer struct {
	ramCap int
	hddCap int
	cpu    string
	os     string
	gpu    string
}

// Each builder should implement this interface 
type PCBuilder interface {
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetGPU() PCBuilder
	GetPC() PersonalComputer
}

// Hoem edition pc builder
type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

// You may ask "Wait? we can set those values on struct initialization.Why
// brand new pattern for it?" and you are right. But remember this is just a
// concise implementation for with informational intend. In reality your
// construction methods will be more sophisticated, so struct initialization 
// will not help.
func (b *HomeEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 4
	return b
}

func (b *HomeEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i3"
	return b
}

func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Home Edition"
	return b
}

func (b *HomeEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "Intel Graphic Card"
	return b
}

func (b *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}

// Game edition pc builder
type GameEditionPCBuilder struct {
	pc PersonalComputer
}

func (b *GameEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 16
	return b
}

func (b *GameEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

func (b *GameEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i7"
	return b
}

func (b *GameEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Ultimate"
	return b
}

func (b *GameEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "AMD Radeon X80"
	return b
}

func (b *GameEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}

//Manufacturer object which aware of build process for builder type
type Manufacturer struct {
	b PCBuilder
}

func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.b = builder
}

func (m *Manufacturer) ConstructPC() {
	m.b.SetCPU().SetHDD().SetRAM().SetGPU().SetOS()
}
