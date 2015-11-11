package thermostat_kata

type HVAC interface {
	Temperature() int

	IsBlowing() bool
	IsCooling() bool
	IsHeating() bool

	SetBlower(state bool)
	SetCooler(state bool)
	SetHeater(state bool)
}
