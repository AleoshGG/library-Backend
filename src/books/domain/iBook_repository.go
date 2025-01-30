package domain

type IBook interface {
	CreateBook()
	GetAllBooks()
	GetBookById()
	GetBookByTitle()
	UpdateBook()
	DeleteBook()
}