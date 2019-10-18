# thermostat-kata

A rendition of Uncle Bob's 
[Environment Controller kata](https://github.com/unclebob/environmentcontroller), 
written in go.

[Watch Presentation](http://university.8thlight.com/events/86) by Uncle Bob


## The scenario

You've been tasked with writing software for a thermostat that will control
HVAC hardware (heater, blower, air-conditioner). This hardware must be operated
according to the manufacturer's specifications in order to run safely and effectively.

To effectively regulate the environment the temperature should never deviate more than
5 degrees from 70F. If the temperature is ever more than 10 degrees away from
70F then the appropriate Hi or Lo alarm must sound.

Bad news:  The various hardware devices have not yet arrived so we can't connect to them yet.
Good news: We have the hardware documentation! The hardware interfaces are defined as follows:

```go
package contracts

type Thermometer interface {
	CurrentTemperature() int // Current ambient temperature rounded to the nearest degree (Fahrenheit).
}

type HVAC interface {
	SetBlower(state bool) // Turns the blower on or off.
	SetCooler(state bool) // Turns the cooler on or off.
	SetHeater(state bool) // Turns the heater on or off.

	IsBlowing() bool // Is the blower currently on or off?
	IsCooling() bool // Is the cooler currently on or off?
	IsHeating() bool // Is the heater currently on or off?
}

type Alarm interface {
	SetColdAlarm(state bool) // Turns the cold alarm on or off.
	SetHeatAlarm(state bool) // Turns the heat alarm on or off.

	ColdAlarm() bool // Is the cold alarm currently on or off?
	HeatAlarm() bool // Is the heat alarm currently on or off?
}
```

On and off are analogous to the go values `true` and `false`, respectively.

## Rules specified by hardware manufacturer:

- When the heater is on, the blower must also be on.
- When the cooler is on, the blower must also be on.
- The blower must continue to run for 5 minutes after the heater is turned off.
    - Because the heat exchanger has latent heat that must not be allowed to
    accumulate lest it melt the sensitive vanes.
- The cooler must not be turned on within 3 minutes of being turned off.
    - Because the freon must be given time to re-condense lest the compressor vapor lock.

---

## Design Questions:

- How will the passage of time be marked?
- How can the thermostat be developed and tested given that the hardware has not yet arrived?

## To Do Items:

Basic Logic:

- TODO: 1. on startup everything is off
- TODO: 2. turn on cooler and blower if too hot
- TODO: 3. nothing turns on if temp just right
- TODO: 4. turn on heater and blower if too cold

Intermediate Logic:

- TODO: 1. cooler turns off if too cold again
- TODO: 2. heater turns off if too hot again
- TODO: 3. blower stays on when cooling after heating
- TODO: 4. blower stays on when heating after cooling

Advanced Logic:

- TODO: 1. heater turns off but blower remains on for 5 minutes after heating
- TODO: 2. cooler must remain off for three minutes before restart
- TODO: 3. cooler does not turn back on after delay if not needed

Alarm Logic:

- TODO: 1. turn on heat alarm at threshold
- TODO: 2. turn on cold alarm at threshold
- TODO: 3. no alarm returned if temp nominal
- TODO: 4. no alarm returned when high temp <= nominal
- TODO: 5. no alarm returned when low temp >= nominal
