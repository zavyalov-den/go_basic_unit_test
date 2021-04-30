package main

import (
	"math"
)

// Incrementor is a struct which contains a number, which is a current number value, a maximum number,
// which is a maximum allowed value for a number. Number can be incremented by a 1 using an incrementValue method
// or could be set to 0 if the current number's value equals to the incrementor's current maximumValue.
type incrementor struct {
	number       uint8
	maximumValue uint8
}

// getNumber returns current number's value.
func (i *incrementor) getNumber() uint8 {
	return i.number
}

// incrementNumber increases the current number value by 1 if current value is less than possible maximum value
// otherwise it sets number value to zero.
func (i *incrementor) incrementNumber() {
	if i.number < math.MaxUint8 {
		i.number++
	} else {
		i.number = 0
	}
}

// setMaximumValue sets incrementor's maximum value to a given value.
// if the new value is higher than current number's value, setMaximumValue sets incrementor's maximum value
// to current number's value.
func (i *incrementor) setMaximumValue(newMaximumValue uint8) {
	if i.number > newMaximumValue {
		i.maximumValue = i.number
	} else {
		i.maximumValue = newMaximumValue
	}
}

// newIncrementor is a helper function that returns a new incrementor struct with default values.
// Current number sets to 0, and maximum value sets to the maximum uint8 value.
func newIncrementor() incrementor {
	return incrementor{
		number:       0,
		maximumValue: math.MaxUint8,
	}
}
