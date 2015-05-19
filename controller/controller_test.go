package controller

//go:generate gunit

/*
Test Cases:
1. On startup everything is off
2a. nothingTurnsOnIfTempJustRight
2b. noAlarmReturnedIfTempNominal
3. turnOnCoolerAndBlowerIfTooHot
4. turnOnHeaterAndBlowerIfTooCold
5. HeatAlarmReturnedAtThreshold
6. ColdAlarmReturnedAtThreshold
7. NoAlarmReturnedWhenHighTempReturnsToNominal
8. NoAlarmReturnedWhenLowTempReturnsToNominal
9. heaterTurnsOffButBlowerRemainsOnFor5MinutesAfterHeating
10. coolerTurnsOffIfHotAgain
11. coolerMustRemainOffForThreeMinutesBeforeRestart
12. coolerDoesNotTurnOnAfterDelayIfNotNeeded
13. blowerDelaysAfterFurnaceGoesOff (necessary?)
14. blowerStaysOnWhenCoolingAfterHeating
*/
