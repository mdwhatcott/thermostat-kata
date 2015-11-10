package controller

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/mdwhatcott/thermostat-kata/hvac"
)

//go:generate gunit

type FakeHardware struct {
	current hvac.Temperature
	blowing bool
	cooling bool
	heating bool
	tooCold bool
	tooHot  bool
}

func NewFakeHardware() *FakeHardware {
	return &FakeHardware{
		blowing: true,
		cooling: true,
		heating: true,
		tooCold: true,
		tooHot:  true,
	}
}

func (this *FakeHardware) Temperature() hvac.Temperature { return this.current }

func (this *FakeHardware) IsHeating() bool { return this.heating }
func (this *FakeHardware) IsCooling() bool { return this.cooling }
func (this *FakeHardware) IsBlowing() bool { return this.blowing }
func (this *FakeHardware) ColdAlarm() bool { return this.tooCold }
func (this *FakeHardware) HeatAlarm() bool { return this.tooHot }

func (this *FakeHardware) SetHeater(on bool)    { this.heating = on }
func (this *FakeHardware) SetCooler(on bool)    { this.cooling = on }
func (this *FakeHardware) SetBlower(on bool)    { this.blowing = on }
func (this *FakeHardware) SetColdAlarm(on bool) { this.tooCold = on }
func (this *FakeHardware) SetHeatAlarm(on bool) { this.tooHot = on }

func (this *FakeHardware) State() string {
	writer := new(bytes.Buffer)
	this.writeState(writer, this.blowing, "BLOWING")
	this.writeState(writer, this.cooling, "COOLING")
	this.writeState(writer, this.heating, "HEATING")
	this.writeState(writer, this.tooCold, "TOO COLD!")
	this.writeState(writer, this.tooHot, "TOO HOT!")
	return strings.TrimSpace(writer.String())
}
func (this *FakeHardware) writeState(writer io.Writer, state bool, message string) {
	if state {
		fmt.Fprint(message + " ")
	}
}

/*
Basic Logic:
TODO: 1. on startup everything is off
TODO: 2. turn on cooler and blower if too hot
TODO: 3. nothing turns on if temp just right
TODO: 4. turn on heater and blower if too cold

Intermediate Logic:
TODO: 1. cooler turns off if too hot again
TODO: 2. heater turns off if too cold again
TODO: 3. blower stays on when cooling after heating
TODO: 4. blower stays on when heating after cooling

Advanced Logic:
TODO: 1. heater turns off but blower remains on for 5 minutes after heating
TODO: 2. cooler must remain off for three minutes before restart
TODO: 3. cooler does not turn back on after delay if not needed

Alarm Logic:
TODO: 1. turn on heat alarm at threshold
TODO: 2. turn on cold alarm at threshold
TODO: 3. no alarm returned if temp nominal
TODO: 4. no alarm returned when high temp <= nominal
TODO: 5. no alarm returned when low temp >= nominal
*/
