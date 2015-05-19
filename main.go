package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/thermostat-kata/controller"
	"github.com/mdwhatcott/thermostat-kata/hvac"
)

func main() {
	hardware := hvac.New()
	controls := controller.New(hardware, 72)

	for {
		controls.Regulate()
		report(hardware)
		time.Sleep(interval)
	}
}

func report(hardware hvac.Hardware) {
	if hardware.ColdAlarm() {
		fmt.Println("[WARN] Too Cold!", hardware.AmbientTemperature())
	} else if hardware.HeatAlarm() {
		fmt.Println("[WARN] Too Hot!", hardware.AmbientTemperature())
	} else {
		fmt.Printf("[INFO] Environment stable; Temperature: %d; Cooling: %t; Blowing: %t; Heating: %t\n",
			hardware.AmbientTemperature(),
			hardware.IsCooling(),
			hardware.IsBlowing(),
			hardware.IsHeating())
	}
}

var interval = time.Minute
