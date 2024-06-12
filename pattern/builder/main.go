package main

type Computer struct {
	CPU     string
	RAM     string
	Storage string
}

type Builder interface {
	SetCPU()
	SetRAM()
	SetStorage()
	GetComputer() Computer
}

type PC struct{ pc Computer }

func (p *PC) SetCPU()               { p.pc.CPU = "Ryzen 3600" }
func (p *PC) SetRAM()               { p.pc.RAM = "16GB" }
func (p *PC) SetStorage()           { p.pc.Storage = "1TB" }
func (p *PC) GetComputer() Computer { return p.pc }

type anotherPC struct{ pc Computer }

func (a *anotherPC) SetCPU()               { a.pc.CPU = "Ryzen 2600" }
func (a *anotherPC) SetRAM()               { a.pc.RAM = "8GB" }
func (a *anotherPC) SetStorage()           { a.pc.Storage = "500GB" }
func (a *anotherPC) GetComputer() Computer { return a.pc }

type Director struct{ builder Builder }

func (d *Director) SetBuilder(builder Builder) { d.builder = builder }

func (d *Director) ConstructComputer() {
	d.builder.SetCPU()
	d.builder.SetRAM()
	d.builder.SetStorage()
}
