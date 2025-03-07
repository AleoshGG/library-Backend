package aplication

import "library-Backend/src/books/domain"

type GetBookByTitle struct {
	db domain.IBook
}

func NewGetBookByTitle(db domain.IBook) *GetBookByTitle {
	return &GetBookByTitle{db: db}
}

func (uc *GetBookByTitle) Run(title string) []domain.Book {
	return uc.db.GetBookByTitle(title)
}