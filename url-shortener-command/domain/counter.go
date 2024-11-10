package domain

type Counter struct {
	currentValue uint64
}

type CounterRepository interface {
	GetNextCounter() (uint64, error)
}
