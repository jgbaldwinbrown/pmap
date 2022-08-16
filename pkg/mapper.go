package pmap

type Mapper[T,U any] struct {
	Input []T
	Func func(T) U
	Threads int
}

func (m Mapper[T,U]) Map() []U {
	return Map(m.Func, m.Input, m.Threads)
}

type Filterer[T any] Mapper[T, bool]

func (f Filterer[T]) Filter() []T {
	return Filter(f.Func, f.Input, f.Threads)
}
