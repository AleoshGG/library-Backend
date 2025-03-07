package validators

import (
	"errors"
	"library-Backend/src/books/domain"
)

func CheckLoan(loan domain.Loan) error {
	if loan.Id_book <= 0 {
		return errors.New("El id debe ser mayor a 0")
	}
	if loan.Id_reader <= 0 {
		return errors.New("El id debe ser mayor a 0")
	}
	if loan.Return_date == "" {
		return errors.New("Verifique la fecha de retorno")
	}
	return nil
}