package hvac

import (
	"fmt"
	"math/rand"
	"strings"
)

// HVAC is a rigged implementation of the Hardware interface
// which chooses a random temperature and rises or falls toward
// the ideal temperature. When that temperature is reached, another
// random temperature is chosen. This allows the controller to
// respond to a somewhat realistic environment.
// It is intended that this will be replaced by the manufacture
// implementation later in development.
type HVAC struct {
	current   Temperature
	heating   bool
	cooling   bool
	blowing   bool
	coldAlarm bool
	heatAlarm bool
}

func New() Hardware {
	return &HVAC{current: 70}
}

func (this *HVAC) Temperature() Temperature {
	if this.current > 70 {
		this.current--
	} else if this.current < 70 {
		this.current++
	} else {
		this.current = Temperature(rand.Intn(77) + 50)
	}
	return this.current
}

func (this *HVAC) SetHeater(on bool)    { this.heating = on }
func (this *HVAC) SetCooler(on bool)    { this.cooling = on }
func (this *HVAC) SetBlower(on bool)    { this.blowing = on }
func (this *HVAC) SetColdAlarm(on bool) { this.coldAlarm = on }
func (this *HVAC) SetHeatAlarm(on bool) { this.heatAlarm = on }

func (this *HVAC) IsHeating() bool { return this.heating }
func (this *HVAC) IsCooling() bool { return this.cooling }
func (this *HVAC) IsBlowing() bool { return this.blowing }
func (this *HVAC) ColdAlarm() bool { return this.coldAlarm }
func (this *HVAC) HeatAlarm() bool { return this.heatAlarm }

func (this *HVAC) String() string {
	state := fmt.Sprintf("%3d degrees; ", this.current)
	state += status(this.IsBlowing(), "BLOWING")
	state += status(this.IsCooling(), "COOLING")
	state += status(this.IsHeating(), "HEATING")
	state += status(this.HeatAlarm(), "TOO HOT!")
	state += status(this.ColdAlarm(), "TOO COLD!")
	return strings.TrimSpace(state)
}
func status(state bool, message string) string {
	if state {
		return message + " "
	} else {
		return strings.Repeat(" ", len(message)+1)
	}
}
