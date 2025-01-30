package infrastructure

import (
	"fmt"
	"library-Backend/src/books/domain"
	"library-Backend/src/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateBook(book domain.Book) (uint, error) {
	query := "INSERT INTO books (title, date_publication, editorial, amount) VALUES (?,?,?,?)"

	res, err := mysql.conn.ExecutePreparedQuery(query, book.Title, book.Date_publication, book.Editorial, book.Amount)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId() 

	return uint(id), nil
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