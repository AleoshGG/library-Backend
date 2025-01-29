package infrastructure

import "fmt"

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (mysql *MySQL) Save() {
	fmt.Println("Producto salvado")
}

func (mysql *MySQL) GetAll() {
	fmt.Println("Lista de productos")
}
