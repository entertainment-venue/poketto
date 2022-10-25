package queue

type Queue interface {
	Offer([]byte) error
	Poll() ([]byte, error)
	Size() int64
}

type OrderedQueue interface {
	Queue
	OfferWithScore([]byte, float64) error
	PollPeak() ([]byte, error)
	PollTail() ([]byte, error)
}
