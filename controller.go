package thermostat_kata

type Controller struct {
	hardware         HVAC
	idealTemperature int
	tolerance        int

	blowerTimer int
	coolerTimer int
}

func NewController(hardware HVAC) *Controller {
	hardware.SetBlower(false)
	hardware.SetCooler(false)
	hardware.SetHeater(false)
	return &Controller{
		hardware: hardware,
	}
}

func (this *Controller) Calibrate(idealTemperature, tolerance int) {
	this.idealTemperature = idealTemperature
	this.tolerance = tolerance
}

func (this *Controller) Regulate() {
	this.decrementTimers()
	temperature := this.hardware.Temperature()

	if this.tooHot(temperature) {
		this.coolIt()
	} else if this.tooCold(temperature) {
		this.heatIt()
	} else {
		this.keepIt()
	}
}

func (this *Controller) decrementTimers() {
	if this.blowerTimer > 0 {
		this.blowerTimer--
	}
	if this.coolerTimer > 0 {
		this.coolerTimer--
	}
}

func (this *Controller) coolIt() {
	this.hardware.SetBlower(true)
	this.hardware.SetCooler(this.coolerTimer == 0)
	this.hardware.SetHeater(false)
}

func (this *Controller) heatIt() {
	this.hardware.SetBlower(true)
	this.hardware.SetCooler(false)
	this.hardware.SetHeater(true)
}

func (this *Controller) keepIt() {
	this.turnBlowerOff()
	this.turnCoolerOff()
	this.turnHeaterOff()
}

func (this *Controller) turnHeaterOff() {
	this.hardware.SetHeater(false)
	if this.blowerTimer == 0 {
		this.blowerTimer = 5
	}
}

func (this *Controller) turnCoolerOff() {
	if this.hardware.IsCooling() {
		this.coolerTimer = 3
	}
	this.hardware.SetCooler(false)

}

func (this *Controller) turnBlowerOff() {
	if this.blowerTimer == 0 {
		this.hardware.SetBlower(this.hardware.IsHeating())
	}
}

func (this *Controller) tooHot(temperature int) bool {
	return temperature > this.idealTemperature+this.tolerance
}
func (this *Controller) tooCold(temperature int) bool {
	return temperature < this.idealTemperature-this.tolerance
}
