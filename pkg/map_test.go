package pmap

import (
	"testing"
)

type Person struct {
	Name string
	Height float64
}

func getPeople() []Person {
	return []Person{Person{"Jim", 74}, Person{"Lolo", 69}}
}

func cmpSlices[T comparable](ps1, ps2 []T) bool {
	if len(ps1) != len(ps2) {
		return false
	}
	for i, p1 := range ps1 {
		if p1 != ps2[i] {
			return false
		}
	}
	return true
}

func TestMap(t *testing.T) {
	people := getPeople()

	heights2 := Map(func(p Person) float64 {return p.Height + 9000}, people, 1000)
	expected := []float64{9074, 9069}
	if !cmpSlices(heights2, expected) {
		t.Errorf("heights2: %v; expected: %v", heights2, expected)
	}

}

func TestFilter(t *testing.T) {
	people := getPeople()
	heights2 := Filter(func(p Person)bool{return p.Height<70}, people, 1000)
	expected := []Person{Person{"Lolo", 69}}
	if !cmpSlices(heights2, expected) {
		t.Errorf("heights2: %v; expected: %v", heights2, expected)
	}
}

func TestReduce(t *testing.T) {
	people := getPeople()
	heights := Map(func(p Person) float64 {return p.Height}, people, -1)
	sum := Reduce(func(f1, f2 float64) float64 {return f1 + f2}, heights, -1)
	if sum != 69 + 74 {
		t.Errorf("reduce: actual: %v; expected: %v", sum, 69+74)
	}
}
