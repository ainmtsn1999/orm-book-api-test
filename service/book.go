package service

import (
	"errors"

	"github.com/ainmtsn1999/orm-book-api-test/model/entity"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(book entity.Book) (entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	GetBook(id int) (entity.Book, error)
	UpdateBook(id int, book entity.Book) (entity.Book, error)
	DeleteBook(id int) error
}

func (s *Service) CreateBook(book entity.Book) (entity.Book, error) {
	// required validation
	if book.NameBook == "" {
		return entity.Book{}, errors.New("name_book_required")
	} else if book.Author == "" {
		return entity.Book{}, errors.New("author_name_required")
	}

	// len validation
	if len(book.NameBook) < 5 {
		return entity.Book{}, errors.New("name_book_less_than_5_letters")
	} else if len(book.Author) < 5 {
		return entity.Book{}, errors.New("author_name_less_than_5_letters")
	}

	newBook := entity.Book{
		NameBook: book.NameBook,
		Author:   book.Author,
	}

	createdBook, err := s.repo.Create(newBook)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Book{}, errors.New("already_exist")
		}
		return entity.Book{}, err
	}

	return createdBook, nil
}

func (s *Service) GetAllBooks() ([]entity.Book, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return []entity.Book{}, errors.New("no_data")
	}
	return books, nil
}

func (s *Service) GetBook(id int) (entity.Book, error) {
	book, err := s.repo.GetOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New("not_found")
		}
		return entity.Book{}, err
	}
	return book, nil
}

func (s *Service) UpdateBook(id int, book entity.Book) (entity.Book, error) {
	// required validation
	if book.NameBook == "" {
		return entity.Book{}, errors.New("name_book_required")
	} else if book.Author == "" {
		return entity.Book{}, errors.New("author_name_required")
	}

	// len validation
	if len(book.NameBook) < 5 {
		return entity.Book{}, errors.New("name_book_less_than_5_letters")
	} else if len(book.Author) < 5 {
		return entity.Book{}, errors.New("author_name_less_than_5_letters")
	}

	// Check kalau ada id
	_, err := s.repo.GetOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New("not_found")
		}
		return entity.Book{}, err
	}

	// Update data
	inputBook := entity.Book{
		NameBook: book.NameBook,
		Author:   book.Author,
	}
	updatedBook, err := s.repo.UpdateOne(id, inputBook)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New("not_found")
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Book{}, errors.New("already_exist")
		}
		return entity.Book{}, err
	}

	return updatedBook, nil
}

func (s *Service) DeleteBook(id int) error {
	// Check kalau ada id
	_, err := s.repo.GetOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not_found")
		}
		return err
	}

	// Lakukan delete
	err = s.repo.DeleteOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("not_found")
		}
		return err
	}
	return nil
}
