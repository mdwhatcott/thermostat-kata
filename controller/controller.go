package controller

import "github.com/mdwhatcott/thermostat-kata/hvac"

type Controller struct {
	hardware   hvac.Hardware
	tooCold    int
	tooHot     int
	wayTooCold int
	wayTooHot  int
}

func New(hardware hvac.Hardware) *Controller {
	hardware.SetHeater(false)
	hardware.SetBlower(false)
	hardware.SetCooler(false)
	return &Controller{hardware: hardware}
}

func (this *Controller) Calibrate(ideal, allowedDelta, alarmDelta int) {
	this.tooCold = ideal - allowedDelta
	this.tooHot = ideal + allowedDelta
	this.wayTooCold = ideal - alarmDelta
	this.wayTooHot = ideal + alarmDelta
}

func (this *Controller) Regulate() {
	current := this.hardware.AmbientTemperature()
	if current > this.tooHot {
		this.hardware.SetCooler(true)
		this.hardware.SetBlower(true)
		this.hardware.SetHeater(false)
		if current > this.wayTooHot {
			this.hardware.SetColdAlarm(false)
			this.hardware.SetHeatAlarm(true)
		}
	} else if current < this.tooCold {
		this.hardware.SetCooler(false)
		this.hardware.SetBlower(true)
		this.hardware.SetHeater(true)
		if current < this.wayTooCold {
			this.hardware.SetColdAlarm(true)
			this.hardware.SetHeatAlarm(false)
		}
	}
}
