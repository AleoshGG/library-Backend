package domain

import "fmt"

type Reader struct {
	Id_reader      int64
	First_name     string
	Last_name      string
	Email          string
	Phone_number   string
	Account_status string
}

func (r *Reader) Show() string {
	return fmt.Sprintf("{id: %d, first_name: %s, last_name: %s, email: %s, phone_number: %s, account_status: %s}",
		r.Id_reader, r.First_name, r.Last_name, r.Email, r.Phone_number, r.Account_status)
}