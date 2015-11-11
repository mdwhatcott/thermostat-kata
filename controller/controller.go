package controller

type Controller struct {
	hardware HVAC
	ideal    int
	delta    int

	blowerOnTimer  int
	coolerOffTimer int
}

func NewController(hardware HVAC) *Controller {
	hardware.SetBlower(false)
	hardware.SetCooler(false)
	hardware.SetHeater(false)
	return &Controller{hardware: hardware}
}

func (this *Controller) SetIdealTemperature(ideal int) { this.ideal = ideal }
func (this *Controller) SetAllowedDelta(delta int)     { this.delta = delta }

func (this *Controller) Regulate() {
	this.decrementTimers()

	if current := this.hardware.Temperature(); this.tooHot(current) {
		this.processTooHot()
	} else if this.justRight(current) {
		this.processJustRight()
	} else if this.tooCold(current) {
		this.processTooCold()
	}
}

func (this *Controller) decrementTimers() {
	if this.blowerOnTimer > 0 {
		this.blowerOnTimer--
	}
	if this.coolerOffTimer > 0 {
		this.coolerOffTimer--
	}
}

func (this *Controller) processTooHot() {
	this.turnOnBlower()
	this.turnOnCooler()
	this.turnOffHeater()
}
func (this *Controller) processJustRight() {
	this.turnOffBlower()
	this.turnOffHeater()
	this.turnOffCooler()
}
func (this *Controller) processTooCold() {
	this.turnOnHeater()
	this.turnOnBlower()
	this.turnOffCooler()
}

func (this *Controller) turnOnHeater() { this.hardware.SetHeater(true) }
func (this *Controller) turnOnBlower() { this.hardware.SetBlower(true) }
func (this *Controller) turnOnCooler() {
	if this.coolerOffTimer == 0 {
		this.hardware.SetCooler(true)
	}
}

func (this *Controller) turnOffHeater() {
	if this.hardware.IsHeating() {
		this.blowerOnTimer = 5
	}
	this.hardware.SetHeater(false)
}
func (this *Controller) turnOffCooler() {
	if this.hardware.IsCooling() {
		this.coolerOffTimer = 3
	}
	this.hardware.SetCooler(false)
}
func (this *Controller) turnOffBlower() {
	if this.blowerOnTimer == 0 {
		this.hardware.SetBlower(this.hardware.IsHeating())
	}
}

func (this *Controller) justRight(current int) bool {
	return !this.tooHot(current) && !this.tooCold(current)
}
func (this *Controller) tooHot(current int) bool {
	return current >= this.ideal+this.delta
}
func (this *Controller) tooCold(current int) bool {
	return current <= this.ideal-this.delta
}
