package aplication

import "library-Backend/src/books/domain"

type GetAllBooks struct {
	db domain.IBook
}

func NewGetAllBooks(db domain.IBook) *GetAllBooks {
	return &GetAllBooks{db: db}
}

func (uc *GetAllBooks) Run() []domain.Book {
	return uc.db.GetAllBooks()
}