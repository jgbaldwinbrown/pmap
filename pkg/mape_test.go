package pmap

import (
	"testing"
)

func TestMapE(t *testing.T) {
	people := getPeople()
	heights2, errs := MapE(func(p Person) (float64, error) {return p.Height + 9000, nil}, people, 1000)
	expected := []float64{9074, 9069}
	if !cmpSlices(heights2, expected) {
		t.Errorf("heights2: %v; expected: %v", heights2, expected)
	}
	for _, err := range errs {
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
