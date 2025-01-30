package domain

type IReader interface {
	CreateReader(reader Reader) (uint, error)
	GetReaderByName(name string) []Reader
	// GetAllReaders()
	// DeleteReader()
	// SetStatusReader()
	// UpdateReader()
}