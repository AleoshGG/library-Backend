package domain

type IBook interface {
	CreateBook(book Book) (uint, error)
	GetAllBooks() []Book
	GetBookById(id_book int) []Book
	GetBookByTitle(title string) []Book
	UpdateBook(id_book int, book Book) (uint, error)
	LendBook(id_book int) (uint, error)
	ReturnBook(id_book int) (uint, error)
	DeleteBook(id_book int) (uint, error)
	GetByEditorialAmount(editorial string, minAmount int) []Book
}