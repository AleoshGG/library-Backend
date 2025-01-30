package aplication

import "library-Backend/src/books/domain"

type CreateBook struct {
	db domain.Book
}

func NewCreateBook(db domain.Book) *CreateBook {
	return &CreateBook{db: db}
}

// Run | Execute
func (uc *CreateBook) Run() {
	uc.db.CreateBook()
}