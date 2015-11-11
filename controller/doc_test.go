package controller

//go:generate gunit

/*
Basic Logic:
TODO: 1. on startup everything is off
TODO: 2. turn on cooler and blower if too hot
TODO: 3. nothing turns on if temp just right
TODO: 4. turn on heater and blower if too cold

Intermediate Logic:
TODO: 1. cooler turns off if too hot again
TODO: 2. heater turns off if too cold again
TODO: 3. blower stays on when cooling after heating
TODO: 4. blower stays on when heating after cooling

Advanced Logic:
TODO: 1. heater turns off but blower remains on for 5 minutes after heating
TODO: 2. cooler must remain off for three minutes before restart
TODO: 3. cooler does not turn back on after delay if not needed

Alarm Logic:
TODO: 1. turn on heat alarm at threshold
TODO: 2. turn on cold alarm at threshold
TODO: 3. no alarm returned if temp nominal
TODO: 4. no alarm returned when high temp <= nominal
TODO: 5. no alarm returned when low temp >= nominal
*/
