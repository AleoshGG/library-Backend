package domain

type IReader interface {
	CreateReader()
	GetReaderByName()
	GetAllReaders()
	DeleteReader()
	SetStatusReader()
	UpdateReader()
}