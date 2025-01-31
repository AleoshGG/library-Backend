package aplication

import "library-Backend/src/readers/domain"

type GetReaderByName struct {
	db domain.IReader
}

func NewGetReaderByName(db domain.IReader) *GetReaderByName {
	return &GetReaderByName{db: db}
}

func (uc *GetReaderByName) Run(name string) []domain.Reader {
	return uc.db.GetReaderByName(name)
}