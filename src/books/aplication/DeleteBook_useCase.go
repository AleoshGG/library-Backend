package aplication

import "library-Backend/src/books/domain"

type DeleteBook struct {
	db domain.IBook
}

func NewDeleteBook(db domain.IBook) *DeleteBook {
	return &DeleteBook{db: db}
}

func (uc *DeleteBook) Run(id_book int) (uint, error) {
	return uc.db.DeleteBook(id_book)
}
