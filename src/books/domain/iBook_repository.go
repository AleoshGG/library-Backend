package domain

type IBook interface {
	CreateBook(book Book) (uint, error)
	GetAllBooks()
	GetBookById()
	GetBookByTitle()
	UpdateBook()
	DeleteBook()
}