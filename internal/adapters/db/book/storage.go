package book

import "Clean-Architecture/internal/domain/book"

type Storage interface {
	GetOne(uuid string) *book.Book
	GetAll(limit, offset int) []book.Book
	Create(book *book.Book) *book.Book
	Delete(uuid string) error
}

func GetOne(uuid string) *book.Book {
	return nil
}

func GetAll(limit, offset int) []book.Book {
	return nil
}

func Create(book *book.Book) *book.Book {
	return nil
}

func Delete(uuid string) error {
	return nil
}
