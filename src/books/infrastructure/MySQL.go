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

func (mysql *MySQL) GetAllBooks() []domain.Book {
	query := "SELECT * FROM books"
	var books []domain.Book
	
	rows := mysql.conn.FetchRows(query)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return books
    }

	defer rows.Close()

	for rows.Next() {
		var b domain.Book
		rows.Scan(&b.Id_book, &b.Title, &b.Date_publication, &b.Editorial, &b.Amount)

		books = append(books, b)
	}
	
	return books  
}

func (mysql *MySQL) GetBookById(id_book int) []domain.Book {
	query := "SELECT * FROM books WHERE id_book = ?"
	var books []domain.Book
	
	rows := mysql.conn.FetchRows(query, id_book)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return books
    }

	defer rows.Close()

	for rows.Next() {
		var b domain.Book
		rows.Scan(&b.Id_book, &b.Title, &b.Date_publication, &b.Editorial, &b.Amount)

		books = append(books, b)
	}
	
	return books
}

func (mysql *MySQL) GetBookByTitle(title string) []domain.Book {
	query := "SELECT * FROM books WHERE title LIKE CONCAT('%', ?, '%')"
	var books []domain.Book
	
	rows := mysql.conn.FetchRows(query, title)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return books
    }

	defer rows.Close()

	for rows.Next() {
		var b domain.Book
		rows.Scan(&b.Id_book, &b.Title, &b.Date_publication, &b.Editorial, &b.Amount)

		books = append(books, b)
	}
	
	return books
}

func (mysql *MySQL) UpdateBook(id_book int, book domain.Book) (uint, error) {
	query := "UPDATE books SET title = ?, date_publication = ?, editorial = ?, amount = ? WHERE id_book = ?" 

	res, err := mysql.conn.ExecutePreparedQuery(query, book.Title, book.Date_publication, book.Editorial, book.Amount, id_book)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}

func (mysql *MySQL) LendBook(id_book int) (uint, error) {
	query := "UPDATE books SET amount = amount - 1 WHERE id_book = ?" 

	res, err := mysql.conn.ExecutePreparedQuery(query,  id_book)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}

func (mysql *MySQL) ReturnBook(id_book int) (uint, error) {
	query := "UPDATE books SET amount = amount + 1 WHERE id_book = ?" 

	res, err := mysql.conn.ExecutePreparedQuery(query,  id_book)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}

func (mysql *MySQL) DeleteBook(id_book int) (uint, error) {
	query := "DELETE FROM books WHERE id_book = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, id_book)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}

func (mysql *MySQL) GetByEditorialAmount(editorial string, minAmount int) []domain.Book {
	query := "SELECT * FROM books WHERE editorial = ? AND amount >= ?"
	var books []domain.Book
	
	rows := mysql.conn.FetchRows(query, editorial, minAmount)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return books
    }

	defer rows.Close()

	for rows.Next() {
		var b domain.Book
		rows.Scan(&b.Id_book, &b.Title, &b.Date_publication, &b.Editorial, &b.Amount)

		books = append(books, b)
	}
	
	return books
}