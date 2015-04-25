package hvac

import (
	"bytes"
	"strings"
)

type HVAC struct{ heating, cooling, blowing bool } /* fake implementation */

func New() *HVAC                  { return &HVAC{} } /* fake implementation */
func (self *HVAC) Heater(on bool) { self.heating = on /* fake implementation */ }
func (self *HVAC) Cooler(on bool) { self.cooling = on /* fake implementation */ }
func (self *HVAC) Blower(on bool) { self.blowing = on /* fake implementation */ }

///////////////////////////////////////////////////////////////////////////////

// String is a debug helper in this case, not a part of the HVAC interface.
func (self *HVAC) String() string {
	buffer := new(bytes.Buffer)
	write(buffer, self.heating, " Heating")
	write(buffer, self.cooling, " Cooling")
	write(buffer, self.blowing, " Blowing")
	return buffer.String()
}
func write(buffer *bytes.Buffer, state bool, message string) {
	if state {
		buffer.WriteString(message)
	} else {
		buffer.WriteString(strings.Repeat(" ", len(message)))
	}
}
