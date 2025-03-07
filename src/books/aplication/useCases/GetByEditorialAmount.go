package aplication

import "library-Backend/src/books/domain"

type GetByEditorialAmount struct {
	db domain.IBook
}

func NewGetByEditorialAmount(db domain.IBook) *GetByEditorialAmount {
	return &GetByEditorialAmount{db: db}
}

func (uc *GetByEditorialAmount) Run(editorial string, minAmount int) []domain.Book {
	return uc.db.GetByEditorialAmount(editorial, minAmount)
}