package thermostat_kata

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type Fixture struct {
	*gunit.Fixture

	hardware   *FakeHardware
	controller *Controller
}

func (this *Fixture) Setup() {
	this.hardware = NewFakeHardware()
	this.controller = NewController(this.hardware)
	this.controller.Calibrate(70, 5)
}

///////////////////////////////////////////////////////////////////////////////

func (this *Fixture) TestUponConstruction_HardwareShouldBeOff() {
	this.assertAllOff()
}

func (this *Fixture) TestWhenTooCold_HeaterAndBlowerTurnOn() {
	this.makeItTooCold()
	this.assertHeating()
}

func (this *Fixture) TestWhenTooHot_CoolerAndBlowerTurnOn() {
	this.makeItTooHot()
	this.assertCooling()
}

func (this *Fixture) TestWhenComfortable_NothingTurnsOn() {
	this.makeItComfy()
	this.assertAllOff()
}

func (this *Fixture) TestHeaterAndCoolerSwitchStates() {
	this.makeItTooCold()
	this.makeItTooHot()
	this.assertCooling()
}

func (this *Fixture) TestCoolerAndHeaterSwitchStates() {
	this.makeItTooHot()
	this.makeItTooCold()
	this.assertHeating()
}

func (this *Fixture) TestAfterCooling_EverythingTurnsOff() {
	this.makeItTooHot()
	this.makeItComfy()
	this.assertAllOff()
}

func (this *Fixture) TestAfterHeating_BlowerStaysOn() {
	this.makeItTooCold()
	this.makeItComfy()
	this.assertBlowing()
}

func (this *Fixture) TestAfterHeating_BlowerStaysOnFor5Minutes() {
	this.makeItTooCold()
	this.makeItComfy()

	for x := 0; x < 5; x++ {
		this.assertBlowing()
		this.makeItComfy()
	}

	this.assertAllOff()
}

func (this *Fixture) TestAfterCooling_CoolerStaysOffFor3MinutesEvenIfNeeded() {
	this.makeItTooHot()
	this.makeItComfy()

	for x := 0; x < 3; x++ {
		this.assertNotCooling()
		this.makeItTooHot()
	}

	this.assertCooling()
}

func (this *Fixture) TestAfterCooling_CoolerStaysOffAfterDelayIfNotNeeded() {
	this.makeItTooHot()

	for x := 0; x < 4; x++ {
		this.makeItComfy()
		this.assertAllOff()
	}
}

///////////////////////////////////////////////////////////////////////////////

func (this *Fixture) makeItTooCold() {
	this.hardware.setTemperature(64)
	this.controller.Regulate()
}
func (this *Fixture) makeItTooHot() {
	this.hardware.setTemperature(76)
	this.controller.Regulate()
}
func (this *Fixture) makeItComfy() {
	this.hardware.setTemperature(70)
	this.controller.Regulate()
}

func (this *Fixture) assertAllOff() {
	this.So(this.hardware.State(), should.Equal, "bch")
}

func (this *Fixture) assertHeating() {
	this.So(this.hardware.State(), should.Equal, "BcH")
}

func (this *Fixture) assertCooling() {
	this.So(this.hardware.State(), should.Equal, "BCh")
}

func (this *Fixture) assertBlowing() {
	this.So(this.hardware.State(), should.Equal, "Bch")
}

func (this *Fixture) assertNotCooling() {
	this.So(this.hardware.State()[1], should.Equal, 'c')
}
