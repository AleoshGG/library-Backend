package aplication

import "library-Backend/src/books/domain"

type GetBookById struct {
	db domain.IBook
}

func NewGetBookById(db domain.IBook) *GetBookById {
	return &GetBookById{db: db}
}

func (uc *GetBookById) Run(id_book int) []domain.Book {
	return uc.db.GetBookById(id_book)
}