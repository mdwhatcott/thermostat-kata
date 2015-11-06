package controller

/*
Basic Logic:
1. on startup everything is off
2. turn on cooler and blower if too hot
3. turn on heater and blower if too cold
4. nothing turns on if temp just right

Advanced Logic:
5. heater turns off but blower remains on for 5 minutes after heating
6. cooler turns off if hot again
7. cooler must remain off for three minutes before restart
8. cooler does not turn back on after delay if not needed
9. blower stays on when cooling after heating

Alarm Logic:
10. turn on heat alarm at threshold
11. turn on cold alarm at threshold
12. no alarm returned if temp nominal
13. no alarm returned when high temp returns to nominal
14. no alarm returned when low temp returns to nominal
*/

/*
type FakeHardware struct {}

func NewFakeHardware() *FakeHardware { return &FakeHardware{} }

func (this *FakeHardware) AmbientTemperature() int { panic("Not implemented") }
func (this *FakeHardware) SetHeater(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetCooler(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetBlower(on bool)       { panic("Not implemented") }
func (this *FakeHardware) SetColdAlarm(on bool)    { panic("Not implemented") }
func (this *FakeHardware) SetHeatAlarm(on bool)    { panic("Not implemented") }
func (this *FakeHardware) IsHeating() bool         { panic("Not implemented") }
func (this *FakeHardware) IsCooling() bool         { panic("Not implemented") }
func (this *FakeHardware) IsBlowing() bool         { panic("Not implemented") }
func (this *FakeHardware) ColdAlarm() bool         { panic("Not implemented") }
func (this *FakeHardware) HeatAlarm() bool         { panic("Not implemented") }
*/
