package aplication

import "library-Backend/src/books/domain"

type ReturnBook struct {
	db domain.IBook
}

func NewReturnBook(db domain.IBook) *ReturnBook {
	return &ReturnBook{db: db}
}

func (uc *ReturnBook) Run(id_book int) (uint, error) {
	return uc.db.ReturnBook(id_book)
}