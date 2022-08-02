package main

import (
	"fmt"
	"sync"
)

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

func FilterFunc(length int, filter func(i int) bool) []bool {
	ok := make([]bool, length)
	MapFunc(length, func(i int) {ok[i] = filter(i)})
	return ok
}

type Person struct {
	Name string
	Height float64
}

type People []Person

type PeopleToHeight struct {
	People []Person
	Heights []float64
}

type ToTallHeight struct{PeopleToHeight}

func (t ToTallHeight) Run(i int) {
	t.Heights[i] = t.People[i].Height + 335
}

func (p People) MapToFloat(f func(Person)float64) []float64 {
	out := make([]float64, len(p))
	run := func(i int) {
		out[i] = f(p[i])
	}
	MapFunc(len(p), run)
	return out
}

func (ps People) Grow(growth_factor float64) []float64 {
	f := func(p Person) float64 {
		return p.Height + growth_factor
	}
	return ps.MapToFloat(f)
}

func (ps People) Collect(oks []bool) (filtered People) {
	for i, ok := range oks {
		if ok {
			filtered = append(filtered, ps[i])
		}
	}
	return
}

func (ps People) Filter(f func(Person)bool) People {
	run := func(i int) bool {
		return f(ps[i])
	}
	return ps.Collect(FilterFunc(len(ps), run))
}

func (ps People) ToTall(height float64) People {
	filter := func(p Person) bool {
		return p.Height > 50
	}
	return ps.Filter(filter)
}

func main() {
	people := People{Person{"Jim", 74}, Person{"Lolo", 69}, Person{"Jones", 19}, Person{"Jupy", 17}}

	heights2 := make([]float64, len(people))
	MapFunc(len(people), func(i int) {heights2[i] = people[i].Height + 9000})
	fmt.Println(heights2)

	fmt.Println(people.Grow(68))

	fmt.Println(people.ToTall(50))

}
