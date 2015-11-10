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
		log.Println(hardware)
		time.Sleep(regulationInterval)
	}
}

var regulationInterval = time.Second // time.Minute

const (
	IDEAL_TEMPERATURE        = 70
	ALLOWED_DELTA_IN_DEGREES = 5
	ALARM_DELTA_IN_DEGREES   = 10
)
