package aplication

import "library-Backend/src/readers/domain"

type DeleteReader struct {
	db domain.IReader
}

func NewDeleteReader(db domain.IReader) *DeleteReader {
	return &DeleteReader{db: db}
}

func (uc *DeleteReader) Run(id_reader int) (uint, error) {
	return uc.db.DeleteReader(id_reader)
}