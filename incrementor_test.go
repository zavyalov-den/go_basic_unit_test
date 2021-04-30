package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func randomUint8(maxValue uint8) uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(int(maxValue)))
}

func TestIncrementor(t *testing.T) {

	// 1. preparing incrementor object to test.
	testIncrementor := newIncrementor()

	// 2. testing getNumber method.
	getNumberResult := testIncrementor.getNumber()

	// 2.1. initial value
	if getNumberResult != testIncrementor.number {
		t.Errorf("getNumber failed, expected %v, got %v", 0, getNumberResult)
	} else {
		t.Logf("getNumber succeeded, expected %v, got %v", 0, getNumberResult)
	}

	// 2.2. random number
	testIncrementor.number = randomUint8(testIncrementor.maximumValue)

	getNumberResult = testIncrementor.getNumber()

	if getNumberResult != testIncrementor.number {
		t.Errorf("getNumber failed, expected %v, got %v", testIncrementor.number, getNumberResult)
	} else {
		t.Logf("getNumber succeeded, expected %v, got %v", testIncrementor.number, getNumberResult)
	}

	// 3. testing setMaximumValue
	// 3.1 basic case
	testIncrementor.number = 0
	newMaximumValue := randomUint8(testIncrementor.maximumValue)
	testIncrementor.setMaximumValue(newMaximumValue)

	if newMaximumValue != testIncrementor.maximumValue {
		t.Errorf("setMaximumValue failed, expected %v, got %v", testIncrementor.maximumValue, newMaximumValue)
	} else {
		t.Logf("setMaximumValue  succeeded, expected %v, got %v", testIncrementor.maximumValue, newMaximumValue)
	}
	// 3.2. testing case if current number is higher than a given maxValue.
	testIncrementor.maximumValue = math.MaxUint8
	testIncrementor.number = testIncrementor.maximumValue / 2

	testIncrementor.setMaximumValue(testIncrementor.number - 1)

	if testIncrementor.maximumValue != testIncrementor.number {
		t.Errorf("setMaximumValue failed, expected %v, got %v", testIncrementor.number, testIncrementor.maximumValue)
	} else {
		t.Logf("setMaximumValue  succeeded, expected %v, got %v", testIncrementor.number, testIncrementor.maximumValue)
	}

	// 4. testing incrementNumber method.
	// 4.1. lower boundary value
	testIncrementor.number = 0
	unchangedNumber := testIncrementor.number

	testIncrementor.incrementNumber()
	incrementedValue := testIncrementor.number

	if unchangedNumber+1 != 1 {
		t.Errorf("incrementValue failed, expected %v, got %v", 1, incrementedValue)
	} else {
		t.Logf("incrementValue succeeded, expected %v, got %v", 1, incrementedValue)
	}

	// 4.2. middle value
	testIncrementor.setMaximumValue(math.MaxUint8)
	unchangedNumber = 127
	testIncrementor.number = unchangedNumber

	testIncrementor.incrementNumber()
	incrementedValue = testIncrementor.number

	if unchangedNumber+1 != incrementedValue {
		t.Errorf("incrementValue failed, expected %v, got %v", unchangedNumber+1, incrementedValue)
	} else {
		t.Logf("incrementValue succeeded, expected %v, got %v", unchangedNumber+1, incrementedValue)
	}

	// 4.3. upper boundary value
	unchangedNumber = testIncrementor.maximumValue
	testIncrementor.number = unchangedNumber
	testIncrementor.incrementNumber()

	if testIncrementor.number != 0 {
		t.Errorf("incrementValue failed, expected %v, got %v", 0, testIncrementor.number)
	} else {
		t.Logf("incrementValue succeeded, expected %v, got %v", 0, testIncrementor.number)
	}

}
