package aplication

import "library-Backend/src/readers/domain"

type UpdateReader struct {
	db domain.IReader
}

func NewUpdateReader(db domain.IReader) *UpdateReader {
	return &UpdateReader{db: db}
}

func (uc *UpdateReader) Run(id_reader int, reader domain.Reader) (uint, error) {
	return uc.db.UpdateReader(id_reader, reader)
}