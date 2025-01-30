package infrastructure

import "fmt"

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (mysql *MySQL) CreateBook() {
	fmt.Println("Producto salvado")
}

func (mysql *MySQL) GetAllBooks() {
	fmt.Println("Lista de productos")
}

func (mysql *MySQL) GetBookById() {
	fmt.Println("Lista de productos")
}

func (mysql *MySQL) GetBookByTitle() {
	fmt.Println("Lista de productos")
}

func (mysql *MySQL) UpdateBook() {
	fmt.Println("Lista de productos")
}

func (mysql *MySQL) DeleteBook() {
	fmt.Println("Lista de productos")
}