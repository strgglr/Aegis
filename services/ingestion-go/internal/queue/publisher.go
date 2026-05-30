package queue

type Publisher interface {
	Publish(event interface{}) error
}

type redisPublisher struct {
}

func (rp *redisPublisher) Publish(event interface{}) error {
	// Implement Redis publishing logic here
	return nil
}

func NewRedisPublisher() Publisher {
	return &redisPublisher{}
}
