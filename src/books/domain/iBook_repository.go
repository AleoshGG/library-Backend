package domain

type IBook interface {
	CreateBook(book Book) (uint, error)
	GetAllBooks() []Book
	GetBookById(id_book int) []Book
	GetBookByTitle(title string) []Book
	UpdateBook(id_book int, book Book) (uint, error)
	DeleteBook(id_book int) (uint, error)
}