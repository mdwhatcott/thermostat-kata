package main

import (
	"log"
	"time"

	"github.com/mdwhatcott/thermostat-kata/controller"
	"github.com/mdwhatcott/thermostat-kata/hvac"
)

func main() {
	hardware := hvac.New()
	thermostat := controller.New(hardware)
	thermostat.Calibrate(IDEAL_TEMPERATURE, ALLOWED_DELTA_IN_DEGREES, ALARM_DELTA_IN_DEGREES)

	for {
		thermostat.Regulate()
		report(hardware)
		time.Sleep(regulationInterval)
	}
}

func report(hardware hvac.Hardware) {
	if hardware.ColdAlarm() {
		log.Println("[WARN] Too Cold!", hardware.AmbientTemperature())
	} else if hardware.HeatAlarm() {
		log.Println("[WARN] Too Hot!", hardware.AmbientTemperature())
	} else {
		log.Printf("[INFO] Environment stable; Temperature: %d; Cooling: %t; Blowing: %t; Heating: %t\n",
			hardware.AmbientTemperature(),
			hardware.IsCooling(),
			hardware.IsBlowing(),
			hardware.IsHeating())
	}
}

var regulationInterval = time.Minute

const (
	IDEAL_TEMPERATURE        = 70
	ALLOWED_DELTA_IN_DEGREES = 5
	ALARM_DELTA_IN_DEGREES   = 10
)
