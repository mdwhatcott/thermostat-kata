package main

import (
	"time"

	"github.com/mdwhatcott/thermostat-kata/controller"
	"github.com/mdwhatcott/thermostat-kata/hvac"
	"github.com/mdwhatcott/thermostat-kata/thermometer"
)

func main() {
	sensor := thermometer.New()
	hardware := hvac.New()
	controls := controller.New(sensor, hardware, 72)

	for {
		controls.Regulate()
		time.Sleep(time.Second)
	}
}
