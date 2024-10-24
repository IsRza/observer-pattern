package util

type Bag struct {
	cancellable []Cancellable
}

func NewBag() *Bag {
	return &Bag{}
}

func (b *Bag) Add(c Cancellable) {
	b.cancellable = append(b.cancellable, c)
}

func (b *Bag) Clean() {
	for _, c := range b.cancellable {
		c.Cancel()
	}
	b.cancellable = nil
}
