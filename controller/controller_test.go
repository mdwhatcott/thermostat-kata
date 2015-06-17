package controller

import (
	"bytes"
	"strings"

	"github.com/smartystreets/gunit"
)

//go:generate gunit

type Fixture struct {
	*gunit.Fixture
}

/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////

/*
Test Cases:
1. On startup everything is off
2a. nothingTurnsOnIfTempJustRight
2b. noAlarmReturnedIfTempNominal
3. turnOnCoolerAndBlowerIfTooHot
4. turnOnHeaterAndBlowerIfTooCold
5. turnOnHeatAlarmAtThreshold
6. turnOnColdAlarmAtThreshold
7. NoAlarmReturnedWhenHighTempReturnsToNominal
8. NoAlarmReturnedWhenLowTempReturnsToNominal
9. heaterTurnsOffButBlowerRemainsOnFor5MinutesAfterHeating
10. coolerTurnsOffIfHotAgain
11. coolerMustRemainOffForThreeMinutesBeforeRestart
12. coolerDoesNotTurnOnAfterDelayIfNotNeeded
13. blowerDelaysAfterFurnaceGoesOff (necessary?)
14. blowerStaysOnWhenCoolingAfterHeating
*/

type FakeHardware struct {
	current   int
	heating   bool
	cooling   bool
	blowing   bool
	coldAlarm bool
	heatAlarm bool
}

func new_fake_hardware() *FakeHardware {
	return &FakeHardware{
		current:   -1,
		heating:   true,
		cooling:   true,
		blowing:   true,
		coldAlarm: true,
		heatAlarm: true,
	}
}

func (this *FakeHardware) AmbientTemperature() int { return this.current }
func (this *FakeHardware) SetHeater(on bool)       { this.heating = on }
func (this *FakeHardware) SetCooler(on bool)       { this.cooling = on }
func (this *FakeHardware) SetBlower(on bool)       { this.blowing = on }
func (this *FakeHardware) SetColdAlarm(on bool)    { this.coldAlarm = on }
func (this *FakeHardware) SetHeatAlarm(on bool)    { this.heatAlarm = on }
func (this *FakeHardware) IsHeating() bool         { return false }
func (this *FakeHardware) IsCooling() bool         { return false }
func (this *FakeHardware) IsBlowing() bool         { return false }
func (this *FakeHardware) ColdAlarm() bool         { return false }
func (this *FakeHardware) HeatAlarm() bool         { return false }

func (this *FakeHardware) State() string {
	buffer := &bytes.Buffer{}
	write(buffer, this.heating, "HEATING")
	write(buffer, this.cooling, "COOLING")
	write(buffer, this.blowing, "BLOWING")
	write(buffer, this.coldAlarm, "COLD!")
	write(buffer, this.heatAlarm, "HOT!")
	return strings.TrimSpace(buffer.String())
}
func write(writer *bytes.Buffer, condition bool, message string) {
	if condition {
		writer.WriteString(message)
		writer.WriteString(" ")
	}
}
