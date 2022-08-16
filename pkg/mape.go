package pmap

type Result[T any] struct {
	Value T
	Err error
}

func Ewrap[T, U any](f func(T) (U, error)) func(T) Result[U] {
	return func(t T) Result[U] {
		var r Result[U]
		r.Value, r.Err = f(t)
		return r
	}
}

func MapE[T, U any](f func(T) (U, error), input []T, threads int) ([]U, []error) {
	results := Map(Ewrap(f), input, threads)
	out := make([]U, len(results))
	errs := make([]error, len(results))
	for i, res := range results {
		out[i] = res.Value
		errs[i] = res.Err
	}
	return out, errs
}

func ResultValues[T any](results []Result[T]) []T {
	out := make([]T, len(results))
	for i, res := range results {
		out[i] = res.Value
	}
	return out
}

func ResultErrs[T any](results []Result[T]) []error {
	out := make([]error, len(results))
	for i, res := range results {
		out[i] = res.Err
	}
	return out
}
