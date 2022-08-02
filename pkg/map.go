package pmap

import (
	"fmt"
	"sync"
)

func Map[T, U any](f func(T)U, input []T, threads int) []U {
	var wg sync.WaitGroup

	njobs := len(input)
	wg.Add(njobs)
	out := make([]U, njobs)
	jobq := make(chan int)

	if threads == -1 {
		threads = njobs
	}
	if threads < 1 {
		threads = 1
	}
	for i:=0; i<threads; i++ {
		go func() {
			for idx := range jobq {
				out[idx] = f(input[idx])
				wg.Done()
			}
		}()
	}

	for i:=0; i<njobs; i++ {
		jobq <- i
	}
	close(jobq)

	wg.Wait()
	return out
}

func Filter[T any](f func(T) bool, input []T, threads int) []T {
	oks := Map(f, input, threads)
	var out []T
	for i, ok := range oks {
		if ok {
			out = append(out, input[i])
		}
	}
	return out
}

func Reduce[T any](f func(T, T) T, input []T, threads int) T {
	var out T
	if len(input) > 0 {
		out = input[0]
	}
	for _, val := range input {
		out = f(out, val)
	}
	return out
}

type Person struct {
	Name string
	Height float64
}

type People []Person

func (p People) MapToFloat(f func(Person)float64) []float64 {
	run := func(per Person) float64 {
		return per.Height
	}
	return Map(run, p, 100)
}

func main() {
	people := People{Person{"Jim", 74}, Person{"Lolo", 69}, Person{"Jones", 19}, Person{"Jupy", 17}}

	heights2 := Map(func(p Person) float64 {return p.Height + 9000}, people, 1000)
	fmt.Println(heights2)

	fmt.Println(people.MapToFloat(func(p Person)float64{return p.Height+48}))

	fmt.Println(Filter(func(p Person)bool{return p.Height>50}, people, 1000))

}
