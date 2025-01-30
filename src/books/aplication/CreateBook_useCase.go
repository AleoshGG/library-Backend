package aplication

import "library-Backend/src/books/domain"

type CreateBook struct {
	db domain.IBook
}

func NewCreateBook(db domain.IBook) *CreateBook {
	return &CreateBook{db: db}
}

func (uc *CreateBook) Run(book domain.Book) (uint, error) {
	return uc.db.CreateBook(book)
}