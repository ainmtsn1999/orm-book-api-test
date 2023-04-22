package repository

import (
	"github.com/ainmtsn1999/orm-book-api-test/model/entity"
)

type BookRepository interface {
	Create(book entity.Book) (entity.Book, error)
	GetAll() ([]entity.Book, error)
	GetOne(id int) (entity.Book, error)
	UpdateOne(id int, book entity.Book) (entity.Book, error)
	DeleteOne(id int) error
}

func (r *Repo) Create(book entity.Book) (entity.Book, error) {
	var createdBook entity.Book

	err := r.gorm.Create(&book).Scan(&createdBook).Error
	if err != nil {
		return entity.Book{}, err
	}

	return createdBook, nil
}

func (r *Repo) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.gorm.Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r *Repo) GetOne(id int) (entity.Book, error) {
	var book entity.Book
	err := r.gorm.Take(&book, "id = ?", id).Error
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (r *Repo) UpdateOne(id int, book entity.Book) (entity.Book, error) {
	var updatedBook entity.Book

	err := r.gorm.Model(&updatedBook).Where("id = ?", id).Updates(&book).Scan(&updatedBook).Error
	if err != nil {
		return entity.Book{}, err
	}

	return updatedBook, nil
}

func (r *Repo) DeleteOne(id int) error {
	err := r.gorm.Where("id = ?", id).Delete(&entity.Book{}).Error
	if err != nil {
		return err
	}
	return nil
}
