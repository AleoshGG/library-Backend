package services

import "library-Backend/src/books/aplication/repository"

type NotifyOfReturnEvent struct {
	rmq repository.IRabbit
}

func NewNotifyOfReturnEvent(rmq repository.IRabbit) *NotifyOfReturnEvent {
	return &NotifyOfReturnEvent{rmq: rmq}
}

func (s *NotifyOfReturnEvent) Run() {
	s.rmq.NotifyOfReturn()
}