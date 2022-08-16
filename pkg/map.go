package pmap

import (
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
	for _, val := range input[1:] {
		out = f(out, val)
	}
	return out
}
