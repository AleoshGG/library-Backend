package aplication

import "library-Backend/src/readers/domain"

type SetStatusReader struct {
	db domain.IReader
}

func NewSetStatusReader(db domain.IReader) *SetStatusReader {
	return &SetStatusReader{db: db}
}

func (uc *SetStatusReader) Run(id_reader int, status string) (uint, error) {
	return uc.db.SetStatusReader(id_reader, status)
}