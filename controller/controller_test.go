package controller

//go:generate gunit

/*
1. On startup everything is off
2. nothingTurnsOnIfTempJustRight
3. turnOnCoolerAndBlowerIfTooHot
4. turnOnHeaterAndBlowerIfTooCold
5. turnOnHiTempAlarmAtThreshold
6. turnOnLoTempAlarmAtThreshold
7. hiTempAlarmResetsWhenTempGoesDown
8. loTempAlarmResetsWhenTempGoesUp
9. heaterTurnsOffButBlowerRemainsOnAfterHeating
10. coolerTurnsOffIfHotAgain
11. coolerMustRemainOffForFiveMinBeforeRestart
12. coolerDoesNotTurnOnAfterDelayIfNotNeeded
13. blowerDelaysAfterFurnaceGoesOff
14. blowerStaysOnWhenCoolingAfterHeating
*/
