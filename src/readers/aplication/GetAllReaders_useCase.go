package aplication

import "library-Backend/src/readers/domain"

type GetAllReaders struct {
	db domain.IReader
}

func NewGetAllReaders(db domain.IReader) *GetAllReaders {
	return &GetAllReaders{db: db}
}

func (uc *GetAllReaders) Run() []domain.Reader {
	return uc.db.GetAllReaders()
}