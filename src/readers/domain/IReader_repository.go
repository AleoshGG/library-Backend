package domain

type IReader interface {
	CreateReader(reader Reader) (uint, error)
	GetReaderByName(name string) []Reader
	GetAllReaders() []Reader
	DeleteReader(id_reader int) (uint, error)
	SetStatusReader(id_reader int, status string) (uint, error)
	UpdateReader(id_reader int, reader Reader) (uint, error)
}