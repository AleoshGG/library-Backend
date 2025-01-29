package aplication

import "server01/domain"

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

// Run | Execute
func (uc *CreateUser) Run() {
	uc.db.CreateUser()
}