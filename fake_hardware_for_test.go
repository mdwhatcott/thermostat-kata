package thermostat_kata

import "strings"

type FakeHardware struct {
	temperature int
	blowing     bool
	cooling     bool
	heating     bool
}

func NewFakeHardware() *FakeHardware {
	return &FakeHardware{
		blowing: true,
		cooling: true,
		heating: true,
	}
}

func (this *FakeHardware) Temperature() int { return this.temperature }

func (this *FakeHardware) IsBlowing() bool { return this.blowing }
func (this *FakeHardware) IsCooling() bool { return this.cooling }
func (this *FakeHardware) IsHeating() bool { return this.heating }

func (this *FakeHardware) SetBlower(state bool) { this.blowing = state }
func (this *FakeHardware) SetCooler(state bool) { this.cooling = state }
func (this *FakeHardware) SetHeater(state bool) { this.heating = state }

func (this *FakeHardware) State() string {
	state := deriveState(this.IsBlowing(), "b")
	state += deriveState(this.IsCooling(), "c")
	state += deriveState(this.IsHeating(), "h")
	return state
}
func deriveState(on bool, letter string) string {
	if on {
		letter = strings.ToUpper(letter)
	}
	return letter
}

func (this *FakeHardware) setTemperature(value int) { this.temperature = value }
