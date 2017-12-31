package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	firstCounter := GetInstance()

	if firstCounter == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := firstCounter

	count := firstCounter.AddOne()
	if count != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", count)
	}

	secondCounter := GetInstance()
	if secondCounter != expectedCounter {
		t.Error("Expected same counter in secondCounter but got different instance")
	}

	count = secondCounter.AddOne()
	if count != 2 {
		t.Errorf("After calling AddOne in second counter, count must be 2 but was $d\n", count)
	}
}