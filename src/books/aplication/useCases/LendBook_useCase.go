package aplication

import "library-Backend/src/books/domain"

type LendBook struct {
	db domain.IBook
}

func NewLendBook(db domain.IBook) *LendBook {
	return &LendBook{db: db}
}

func (uc *LendBook) Run(id_book int) (uint, error) {
	return uc.db.LendBook(id_book)
}