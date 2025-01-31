package validators

import (
	"errors"
	"library-Backend/src/books/domain"
)

func CheckBook(book domain.Book) error {
	if book.Amount <= 0 {
		return errors.New("La cantidad debe ser mayor a 0")
	}
	if book.Date_publication == "" {
		return errors.New("Verifique la fecha de publicación")
	}
	if book.Title == "" {
		return errors.New("Verifique el título")
	}
	if book.Editorial == "" {
		return errors.New("Verifique la editorial")
	}
	return nil
}