package controller

import (
	"strings"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

//go:generate gunit

type Fixture struct {
	*gunit.Fixture
	hardware   *FakeHardware
	controller *Controller
}

func (this *Fixture) Setup() {
	this.hardware = NewFakeHardware()
	this.hardware.SetHeater(true)
	this.hardware.SetBlower(true)
	this.hardware.SetCooler(true)
	this.controller = New(this.hardware)
	this.controller.Calibrate(70, 5, 10)
}
func (this *Fixture) assertState(expected string) {
	this.So(this.hardware.State(), should.Equal, expected)
}
func (this *Fixture) assertAllOff() {
	this.assertState("")
}

func (this *Fixture) TestAllHardwareOffAtStartup() {
	this.assertAllOff()
}

func (this *Fixture) TestNothingTurnsOnIfTemperatureJustRight() {
	this.hardware.current = 70
	this.controller.Regulate()
	this.assertAllOff()
}

func (this *Fixture) TestTurnOnCoolerAndBlowerIfTooHot() {
	this.hardware.current = 76
	this.controller.Regulate()
	this.assertState("BLOWING COOLING")
}

func (this *Fixture) TestTurnOnHeaterAndBLowerIfTooCold() {
	this.hardware.current = 64
	this.controller.Regulate()
	this.assertState("HEATING BLOWING")
}

func (this *Fixture) TestAlarms() {
	for x := 0; x < 60; x++ {
		this.assertAlarm(x, "COLD!")
	}
	for x := 60; x <= 80; x++ {
		this.assertAlarm(x, "")
	}
	for x := 81; x <= 130; x++ {
		this.assertAlarm(x, "HOT!")
	}
}
func (this *Fixture) assertAlarm(temperature int, expected string) {
	this.Setup()
	this.hardware.current = temperature
	this.controller.Regulate()
	this.So(this.hardware.Alarms(), should.Equal, expected)
}

/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////

type FakeHardware struct {
	current   int
	heating   bool
	cooling   bool
	blowing   bool
	coldAlarm bool
	heatAlarm bool
}

func NewFakeHardware() *FakeHardware { return &FakeHardware{} }

func (this *FakeHardware) AmbientTemperature() int { return this.current }
func (this *FakeHardware) SetHeater(on bool)       { this.heating = on }
func (this *FakeHardware) SetCooler(on bool)       { this.cooling = on }
func (this *FakeHardware) SetBlower(on bool)       { this.blowing = on }
func (this *FakeHardware) SetColdAlarm(on bool)    { this.coldAlarm = on }
func (this *FakeHardware) SetHeatAlarm(on bool)    { this.heatAlarm = on }
func (this *FakeHardware) IsHeating() bool         { return this.heating }
func (this *FakeHardware) IsCooling() bool         { return this.cooling }
func (this *FakeHardware) IsBlowing() bool         { return this.blowing }
func (this *FakeHardware) ColdAlarm() bool         { return this.coldAlarm }
func (this *FakeHardware) HeatAlarm() bool         { return this.heatAlarm }

func (this *FakeHardware) State() string {
	state := ""
	if this.IsHeating() {
		state += "HEATING "
	}
	if this.IsBlowing() {
		state += "BLOWING "
	}
	if this.IsCooling() {
		state += "COOLING "
	}
	return strings.TrimSpace(state)
}
func (this *FakeHardware) Alarms() string {
	state := ""
	if this.HeatAlarm() {
		state += "HOT! "
	}
	if this.ColdAlarm() {
		state += "COLD!"
	}
	return strings.TrimSpace(state)
}

/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////

/*
Test Cases:
1. On startup everything is off
3. turnOnCoolerAndBlowerIfTooHot
4. turnOnHeaterAndBlowerIfTooCold
2a. nothingTurnsOnIfTempJustRight
5. turnOnHeatAlarmAtThreshold
6. turnOnColdAlarmAtThreshold
2b. noAlarmReturnedIfTempNominal
7. NoAlarmReturnedWhenHighTempReturnsToNominal
8. NoAlarmReturnedWhenLowTempReturnsToNominal
9. heaterTurnsOffButBlowerRemainsOnFor5MinutesAfterHeating
10. coolerTurnsOffIfHotAgain
11. coolerMustRemainOffForThreeMinutesBeforeRestart
12. coolerDoesNotTurnOnAfterDelayIfNotNeeded
13. blowerDelaysAfterFurnaceGoesOff (necessary?)
14. blowerStaysOnWhenCoolingAfterHeating
*/

/*
type FakeHardware struct {}

func NewFakeHardware() *FakeHardware { return &FakeHardware{} }

func (this *FakeHardware) AmbientTemperature() int { panic("Not implemented") }
func (this *FakeHardware) SetHeater(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetCooler(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetBlower(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetColdAlarm(on bool)    { panic("Not implemented") }
func (this *FakeHardware) SetHeatAlarm(on bool)    { panic("Not implemented") }
func (this *FakeHardware) IsHeating() bool         { panic("Not implemented") }
func (this *FakeHardware) IsCooling() bool         { panic("Not implemented") }
func (this *FakeHardware) IsBlowing() bool         { panic("Not implemented") }
func (this *FakeHardware) ColdAlarm() bool         { panic("Not implemented") }
func (this *FakeHardware) HeatAlarm() bool         { panic("Not implemented") }
*/
