package main

import (
	"fmt"
	"sync"
)

type Mapping interface {
	Len() int
	Run(i int)
}

func MapFunc(length int, run func(i int)) {
	var wg sync.WaitGroup
	wg.Add(length)
	for i:=0; i<length; i++ {
		go func(pos int) {
			defer wg.Done()
			run(pos)
			fmt.Println("Proc", pos, "done.")
		}(i)
	}
	wg.Wait()
}

func Map(m Mapping) {
	MapFunc(m.Len(), func(i int){m.Run(i)})
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

	heights2 := make([]float64, len(people))
	MapFunc(len(people), func(i int) {heights2[i] = people[i].Height + 9000})
	fmt.Println(heights2)

	heights := make([]float64, len(people))
	Map(ToTallHeight{PeopleToHeight{people, heights}})
	fmt.Println(heights)
}
