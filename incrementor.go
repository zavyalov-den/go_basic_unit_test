package main

import (
	"math"
)

type incrementor struct {
	number       uint8
	maximumValue uint8
}

func (i *incrementor) getNumber() uint8 {
	return i.number
}

func (i *incrementor) incrementNumber() {
	if i.number < math.MaxUint8 {
		i.number++
	} else {
		i.number = 0
	}
}

func (i *incrementor) setMaximumValue(maximumValue uint8) {
	i.maximumValue = maximumValue
}

func newIncrementor() incrementor {
	return incrementor{
		number:       0,
		maximumValue: math.MaxUint8,
	}
}
