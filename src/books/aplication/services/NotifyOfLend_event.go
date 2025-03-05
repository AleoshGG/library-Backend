package services

import "library-Backend/src/books/aplication/repository"

type NotifyOfLendEvent struct {
	rmq repository.IRabbit
}

func NewNotifyOfLend(rmq repository.IRabbit) *NotifyOfLendEvent {
	return &NotifyOfLendEvent{rmq: rmq}
}

func (s *NotifyOfLendEvent) Run() {
	s.rmq.NotifyOfLend()
}