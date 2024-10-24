package util

type (
	Subject[T any] interface {
		Observe(callback func(T)) Observer[T]
		SetState(state T)
		GetState() T

		detach(id string)
	}

	concreteSubject[T any] struct {
		state     T
		observers map[string]Observer[T]
	}
)

func NewConcreteSubject[T any](initial T) Subject[T] {
	return &concreteSubject[T]{
		state:     initial,
		observers: map[string]Observer[T]{},
	}
}

func (s *concreteSubject[T]) Observe(callback func(T)) Observer[T] {
	o := NewObserver(callback, s)
	s.observers[o.getID()] = o
	callback(s.state)
	return o
}

func (s *concreteSubject[T]) SetState(state T) {
	s.state = state
	for _, o := range s.observers {
		//go o.notify()
		go o.notifyWithState(s.state)
	}
}

func (s *concreteSubject[T]) GetState() T {
	return s.state
}

func (s *concreteSubject[T]) detach(id string) {
	delete(s.observers, id)
}
