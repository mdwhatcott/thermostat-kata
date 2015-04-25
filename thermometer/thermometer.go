package thermometer

import "math/rand"

func New() *Thermometer { return &Thermometer{} }

type Thermometer struct{}

func (self *Thermometer) Measure() int {
	return rand.Intn(80) + 50 /* not a real implementation... */
}
