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
	Filter(i int) bool
	Collect(oks []bool)
}

func FilterFunc(length int, filter func(i int) bool) []bool {
	ok := make([]bool, length)
	MapFunc(length, func(i int) {ok[i] = filter(i)})
	return ok
}

func Filter(f Filterable) {
	oks := FilterFunc(f.Len(), func (i int) bool {return f.Filter(i)})
	f.Collect(oks)
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
	Filtered *[]Person
}

func (f PeopleFilter) Len() int {
	return len(f.People)
}

func (f PeopleFilter) Collect(oks []bool) {
	for i, ok := range oks {
		if ok {
			*f.Filtered = append(*f.Filtered, f.People[i])
		}
	}
}

type ToTallPeople struct {PeopleFilter}

func (t ToTallPeople) Filter(i int) bool {
	return t.People[i].Height > 50
}

func main() {
	people := []Person{Person{"Jim", 74}, Person{"Lolo", 69}, Person{"Jones", 19}, Person{"Jupy", 17}}

	heights2 := make([]float64, len(people))
	MapFunc(len(people), func(i int) {heights2[i] = people[i].Height + 9000})
	fmt.Println(heights2)

	heights := make([]float64, len(people))
	Map(ToTallHeight{PeopleToHeight{people, heights}})
	fmt.Println(heights)

	talls := []Person{}
	Filter(ToTallPeople{PeopleFilter{people, &talls}})
	fmt.Println(talls)

}
