package controller

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type Fixture struct {
	*gunit.Fixture

	hardware   *FakeHVAC
	controller *Controller
}

func (this *Fixture) Setup() {
	this.hardware = NewFakeHVAC()
	this.hardware.SetBlower(true)
	this.hardware.SetCooler(true)
	this.hardware.SetHeater(true)
	this.controller = NewController(this.hardware)
	this.controller.SetIdealTemperature(75)
	this.controller.SetAllowedDelta(3)
}

func (this *Fixture) TestUponConstruction_HardwareDeactivated() {
	this.assertAllOff()
}
func (this *Fixture) TestWhenTooHot_TurnOnBlowerAndCooler() {
	this.makeItTooHot()
	this.assertCooling()
}
func (this *Fixture) TestWhenTooCold_TurnOnBlowerAndHeater() {
	this.makeItTooCold()
	this.assertHeating()
}
func (this *Fixture) TestWhenComfortable_NothingComesOn() {
	this.makeItComfortable()
	this.assertAllOff()
}
func (this *Fixture) TestHeaterNeverRunsWhenTooCold() {
	this.makeItTooColdThenTooHot()
	this.assertCooling()
}
func (this *Fixture) TestCoolerNeverRunsWhenTooHot() {
	this.makeItTooHotThenTooCold()
	this.assertHeating()
}
func (this *Fixture) TestBlowerRemainsOnFor5MinutesWhenHeaterTurnsOff() {
	this.makeItTooCold()
	this.makeItComfortable()

	for x := 0; x < 5; x++ {
		this.assertBlowing()
		this.makeItComfortable()
	}
	this.assertAllOff()
}
func (this *Fixture) TestCoolerStaysOffFor3MinutesAfterShuttingOffEvenWhenNeeded() {
	this.makeItTooHot()
	this.makeItComfortable()

	for x := 0; x < 3; x++ {
		this.assertNotCooling()
		this.makeItTooHot()
	}
	this.assertCooling()
}
func (this *Fixture) TestCoolerDoesNotTurnOnAfterDelayIfNotNeeded() {
	this.makeItTooHot()
	this.makeItComfortable()

	for x := 0; x < 3; x++ {
		this.assertNotCooling()
		this.makeItComfortable()
	}
	this.assertNotCooling()
}

///////////////////////////////////////////////////////////////////////////////

func (this *Fixture) makeItComfortable() {
	this.hardware.setCurrentTemperature(75)
	this.controller.Regulate()
}
func (this *Fixture) makeItTooHot() {
	this.hardware.setCurrentTemperature(78)
	this.controller.Regulate()
}
func (this *Fixture) makeItTooCold() {
	this.hardware.setCurrentTemperature(72)
	this.controller.Regulate()
}
func (this *Fixture) makeItTooColdThenTooHot() {
	this.makeItTooCold()
	this.makeItTooHot()
}
func (this *Fixture) makeItTooHotThenTooCold() {
	this.makeItTooHot()
	this.makeItTooCold()
}

func (this *Fixture) assertAllOff()     { this.So(this.hardware.String(), should.Equal, "bch") }
func (this *Fixture) assertCooling()    { this.So(this.hardware.String(), should.Equal, "BCh") }
func (this *Fixture) assertHeating()    { this.So(this.hardware.String(), should.Equal, "BcH") }
func (this *Fixture) assertBlowing()    { this.So(this.hardware.String(), should.Equal, "Bch") }
func (this *Fixture) assertNotCooling() { this.So(this.hardware.String()[1], should.Equal, 'c') }
