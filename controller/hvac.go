package controller

type HVAC interface {
	Temperature() int

	IsBlowing() bool
	IsCooling() bool
	IsHeating() bool

	SetBlower(on bool)
	SetCooler(on bool)
	SetHeater(on bool)
}

