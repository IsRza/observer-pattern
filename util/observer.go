package util

import "github.com/google/uuid"

type (
	Cancellable interface {
		Cancel()
	}

	Observer[T any] interface {
		Cancellable

		StoreIn(bag *Bag)

		notify()
		notifyWithState(T)
		getID() string
	}

	observer[T any] struct {
		id       string
		callback func(T)
		subject  Subject[T]
	}
)

func NewObserver[T any](callback func(T), subject Subject[T]) Observer[T] {
	return &observer[T]{
		id:       uuid.NewString(),
		callback: callback,
		subject:  subject,
	}
}

func (o *observer[T]) Cancel() {
	o.subject.detach(o.id)
	o.subject = nil
	o.callback = nil
}

func (o *observer[T]) StoreIn(bag *Bag) {
	bag.Add(o)
}

func (o *observer[T]) notify() {
	o.callback(o.subject.GetState())
}

func (o *observer[T]) notifyWithState(state T) {
	o.callback(state)
}

func (o *observer[T]) getID() string {
	return o.id
}
