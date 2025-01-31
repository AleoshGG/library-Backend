package validators

import (
	"errors"
	"library-Backend/src/readers/domain"
)

func CheckReader(reader domain.Reader) error {
	if reader.First_name == "" {
		return errors.New("Verifique el primer nombre")
	}
	if reader.Last_name == "" {
		return errors.New("Verifique los apellidos")
	}
	if reader.Email == "" {
		return errors.New("Verifique el email")
	}
	if reader.Phone_number == "" {
		return errors.New("Verifique el telefono")
	}
	return nil
}