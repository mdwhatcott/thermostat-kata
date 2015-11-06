# thermostat-kata

A rendition of Uncle Bob's [Environment Controller kata](https://github.com/unclebob/environmentcontroller), written in go.


## The scenario

You've been tasked with writing software for a thermostat that will control
HVAC hardware (heater, blower, air-conditioner). This hardware must be operated
according to the manufacturer's specifications in order to run safely and effectively.

In this scenario the temperature must never deviate more than 5 degrees from 70F.
If the temperature is ever more than 10 degrees away from 70F then the appropriate
Hi or Lo alarm must sound.


The hardware interface is as follows:

```go
type Hardware interface {
	AmbientTemperature() int // Returns the current ambient temperature.

	SetHeater(on bool)    // Turns the heater on or off.
	SetCooler(on bool)    // Turns the air conditioner on or off.
	SetBlower(on bool)    // Turns the blower on or off.
	SetColdAlarm(on bool) // Turns the cold alarm on or off.
	SetHeatAlarm(on bool) // Turns the heat alarm on or off.

	IsHeating() bool // Returns whether the heater is currently on or off.
	IsCooling() bool // Returns whether the air conditioner is currently on or off.
	IsBlowing() bool // Returns whether the blower is currently on or off.
	ColdAlarm() bool // Returns whether the current temperature is below the ideal minus the allowed delta.
	HeatAlarm() bool // Returns whether the current temperature is above the ideal plus the allowed delta.
}
```

## Rules specified by hardware manufacturer:

- When the heater is on, the blower must also be on.
- When the cooler is on, the blower must also be on.
- The blower must continue to run for 5 minutes after the heater is turned off.
    - Because the heat exchanger has latent heat that must not be allowed to
    accumulate lest it melt the sensitive vanes.
- The cooler must not be turned on within 3 minutes of being turned off.
    - Because the freon must be given time to re-condense lest the compressor vapor lock.

## Design Questions:

- How will the passage of time be marked?
- How can the thermostat be developed and tested given that the hardware has not yet arrived?
