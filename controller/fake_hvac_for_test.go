package controller
import "strings"

type FakeHVAC struct {
	current int
	blowing bool
	cooling bool
	heating bool
}

func NewFakeHVAC() *FakeHVAC {
	return &FakeHVAC{}
}

func (this *FakeHVAC) setCurrentTemperature(current int) { this.current = current }

func (this *FakeHVAC) Temperature() int { return this.current }

func (this *FakeHVAC) IsBlowing() bool { return this.blowing }
func (this *FakeHVAC) IsCooling() bool { return this.cooling }
func (this *FakeHVAC) IsHeating() bool { return this.heating }

func (this *FakeHVAC) SetBlower(on bool) { this.blowing = on }
func (this *FakeHVAC) SetCooler(on bool ) { this.cooling = on }
func (this *FakeHVAC) SetHeater(on bool) { this.heating = on }

func (this *FakeHVAC) String() string {
	state := stateChar(this.IsBlowing(), "b")
	state += stateChar(this.IsCooling(), "c")
	state += stateChar(this.IsHeating(), "h")
	return state
}
func stateChar(state bool, char string) string {
	if state {
		char = strings.ToUpper(char)
	}
	return char
}