package controller

import "github.com/mdwhatcott/thermostat-kata/hvac"

type FakeHardware struct {
	current hvac.Temperature
	blowing bool
	cooling bool
	heating bool
	tooCold bool
	tooHot  bool
}

func NewFakeHardware() *FakeHardware {
	return &FakeHardware{}
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
