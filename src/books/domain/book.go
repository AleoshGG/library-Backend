package domain

import (
	"fmt"
	"time"
)

type Book struct {
	Id_book          int64
	Title            string
	Date_publication time.Time
	Editorial        string
	Amount           int64
}

func NewBook(title string, date_publication time.Time, editorial string, amount int64) *Book {
	return &Book{Title: title, Date_publication: date_publication, Editorial: editorial, Amount: amount}
}

func (b *Book) Show() string {
	return fmt.Sprintf("{id: %d, title: %s, date_publication: %s, editorial: %s, amount: %d}",
		b.Id_book, b.Title, b.Date_publication, b.Editorial, b.Amount)
}
