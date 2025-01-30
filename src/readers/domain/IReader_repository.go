package domain

type IReader interface {
	CreateReader(reader Reader) (uint, error)
	// GetReaderByName()
	// GetAllReaders()
	// DeleteReader()
	// SetStatusReader()
	// UpdateReader()
}