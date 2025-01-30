package aplication

import "library-Backend/src/books/domain"

type UpdateBook struct {
	db domain.IBook
}

func NewUpdateBook(db domain.IBook) *UpdateBook {
	return &UpdateBook{db: db}
}

func (uc *UpdateBook) Run (id_book int, book domain.Book) (uint, error) {
	return uc.db.UpdateBook(id_book, book)
}