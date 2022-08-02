package main

import (
	"fmt"
	"sync"
)

type Mapping interface {
	Len() int
	Run(i int)
}

func Map(m Mapping) {
	var wg sync.WaitGroup
	wg.Add(m.Len())
	for i:=0; i<m.Len(); i++ {
		go func(pos int) {
			defer wg.Done()
			m.Run(pos)
			fmt.Println("Proc", pos, "done.")
		}(i)
	}
	wg.Wait()
}

type Filterable interface {
	Len() int
	Run(i int)
	Collect()
}

type Person struct {
	Name string
	Height float64
}

type PeopleToHeight struct {
	People []Person
	Heights []float64
}

func (t PeopleToHeight) Len() int {
	if len(t.People) < len(t.Heights) {
		return len(t.People)
	}
	return len(t.Heights)
}

type ToTallHeight struct{PeopleToHeight}

func (t ToTallHeight) Run(i int) {
	t.Heights[i] = t.People[i].Height + 335
}

type PeopleFilter struct {
	People []Person
	Ok []bool
	Filtered []Person
}

func (f PeopleFilter) Len() int {
	return len(f.People)
}

func (f PeopleFilter) Collect() {
	for i:=0; i<len(f.People); i++ {
		if f.Ok[i] {
			f.Filtered = append(f.Filtered, f.People[i])
		}
	}
}

func main() {
	people := []Person{Person{"Jim", 74}, Person{"Lolo", 69}, Person{"Jones", 19}, Person{"Jupy", 17}}
	heights := make([]float64, len(people))
	Map(ToTallHeight{PeopleToHeight{people, heights}})
	fmt.Println(heights)
}
