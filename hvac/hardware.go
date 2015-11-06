package hvac

// Hardware is a manufacture-specified interface for
// all hardware controlled by the thermostat.
type Hardware interface {
	AmbientTemperature() int

	SetHeater(on bool)
	SetCooler(on bool)
	SetBlower(on bool)
	SetColdAlarm(on bool)
	SetHeatAlarm(on bool)

	IsHeating() bool
	IsCooling() bool
	IsBlowing() bool
	ColdAlarm() bool
	HeatAlarm() bool
}
