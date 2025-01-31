package aplication

import "library-Backend/src/readers/domain"

type CreateReader struct {
	db domain.IReader
}

func NewCreateReader(db domain.IReader) *CreateReader {
	return &CreateReader{db: db}
}

func (uc *CreateReader) Run(reader domain.Reader) (uint, error) {
	return uc.db.CreateReader(reader)
}