package service

import (
	"errors"
	"testing"

	"github.com/ainmtsn1999/orm-book-api-test/model/entity"
	"github.com/ainmtsn1999/orm-book-api-test/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookService_GetAllBooks(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult []entity.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepository)
	}

	var testTable []testCase

	// Test jika sukses
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		expectedResult: []entity.Book{
			{
				ID:       2,
				NameBook: "Inayatulloh's Book",
				Author:   "Inayatulloh",
			},
			{
				ID:       3,
				NameBook: "Muiin's Book",
				Author:   "Muiin",
			},
		},
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetAll().Return([]entity.Book{
				{
					ID:       2,
					NameBook: "Inayatulloh's Book",
					Author:   "Inayatulloh",
				},
				{
					ID:       3,
					NameBook: "Muiin's Book",
					Author:   "Muiin",
				},
			}, nil).Times(1)
		},
	})

	// Test jika tidak ada data buku
	testTable = append(testTable, testCase{
		name:      "no books",
		wantError: true,
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetAll().Return([]entity.Book{}, errors.New("no_data")).Times(1)
		},
		expectedError: errors.New("no_data"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Buat mock controller
			mockController := gomock.NewController(t)
			bookRepo := mocks.NewMockBookRepository(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetAllBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestBookService_GetBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult entity.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepository)
	}

	var testTable []testCase

	// Test jika ada data buku
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		expectedResult: entity.Book{
			ID:       2,
			NameBook: "Inayatulloh's Book",
			Author:   "Inayatulloh",
		},
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{
				ID:       2,
				NameBook: "Inayatulloh's Book",
				Author:   "Inayatulloh",
			}, nil).Times(1)
		},
	})

	// Test jika tidak ada buku
	testTable = append(testTable, testCase{
		name:          "no books",
		wantError:     true,
		expectedError: errors.New("not_found"),
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{}, gorm.ErrRecordNotFound)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Buat mock controller
			mockController := gomock.NewController(t)
			bookRepo := mocks.NewMockBookRepository(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBook(2)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestBookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          entity.Book
		expectedResult entity.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepository)
	}

	var testTable []testCase

	// test success add data
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: entity.Book{
			NameBook: "Muiin's Book",
			Author:   "Muiin",
		},
		expectedResult: entity.Book{
			ID:       1,
			NameBook: "Muiin's Book",
			Author:   "Muiin",
		},
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().Create(gomock.Any()).Return(entity.Book{
				ID:       1,
				NameBook: "Muiin's Book",
				Author:   "Muiin",
			}, nil).Times(1)
		},
	})

	// test error book already exist
	testTable = append(testTable, testCase{
		name:      "book already exist",
		wantError: true,
		input: entity.Book{
			NameBook: "Muiin's Book",
			Author:   "Muiin",
		},
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().Create(gomock.Any()).Return(entity.Book{}, gorm.ErrDuplicatedKey).Times(1)
		},
		expectedError: errors.New("already_exist"),
	})

	// test error required field namebook
	testTable = append(testTable, testCase{
		name:      "validation name_book required",
		wantError: true,
		input: entity.Book{
			Author: "Muiin",
		},
		expectedError: errors.New("name_book_required"),
	})

	// test error required field author
	testTable = append(testTable, testCase{
		name:      "validation author required",
		wantError: true,
		input: entity.Book{
			NameBook: "Muiin's Book",
		},
		expectedError: errors.New("author_name_required"),
	})

	// test error name_book tidak boleh kurang dari 3 huruf
	testTable = append(testTable, testCase{
		name:      "validation name_book less than 5 letters",
		wantError: true,
		input: entity.Book{
			NameBook: "Yo",
			Author:   "Inayatulloh",
		},
		expectedError: errors.New("name_book_less_than_5_letters"),
	})

	// test error author tidak boleh kurang dari 3 huruf
	testTable = append(testTable, testCase{
		name:      "validation author less than 5 letters",
		wantError: true,
		input: entity.Book{
			NameBook: "Muiin's Book",
			Author:   "Da",
		},
		expectedError: errors.New("author_name_less_than_5_letters"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Buat mock controller
			mockController := gomock.NewController(t)
			bookRepo := mocks.NewMockBookRepository(mockController)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			// Ambil hasil
			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestBookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          entity.Book
		idBook         int
		expectedResult entity.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepository)
	}

	var testTable []testCase

	// test success add data
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: entity.Book{
			NameBook: "Inayatulloh's Book",
			Author:   "Inayatullloh",
		},
		idBook: 1,
		expectedResult: entity.Book{
			ID:       1,
			NameBook: "Inayatulloh's Book",
			Author:   "Inayatullloh",
		},
		onBookRepo: func(mock *mocks.MockBookRepository) {
			// Dua mock dikarenakan dalam service menggunakan dua repo
			// GetOne digunakan untuk check id
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{
				ID:       1,
				NameBook: "Muiin's Book",
				Author:   "Muiin",
			}, nil).Times(1)
			mock.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Return(entity.Book{
				ID:       1,
				NameBook: "Inayatulloh's Book",
				Author:   "Inayatullloh",
			}, nil).Times(1)
		},
	})

	// test error book jika tidak ada buku ditemukan
	testTable = append(testTable, testCase{
		name:      "book not exist",
		wantError: true,
		input: entity.Book{
			NameBook: "Inayatulloh's Book",
			Author:   "Inayatullloh",
		},
		idBook: 1,
		onBookRepo: func(mock *mocks.MockBookRepository) {
			// Dua mock dikarenakan dalam service menggunakan dua repo
			// GetOne digunakan untuk check id
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{}, gorm.ErrRecordNotFound).Times(1)
			mock.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Return(entity.Book{}, gorm.ErrRecordNotFound).AnyTimes()
		},
		expectedError: errors.New("not_found"),
	})

	// test error required field namebook
	testTable = append(testTable, testCase{
		name:      "validation name_book required",
		wantError: true,
		idBook:    1,
		input: entity.Book{
			Author: "Muiin",
		},
		expectedError: errors.New("name_book_required"),
	})

	// test error required field author
	testTable = append(testTable, testCase{
		name:      "validation author required",
		wantError: true,
		idBook:    1,
		input: entity.Book{
			NameBook: "Muiin's Book",
		},
		expectedError: errors.New("author_name_required"),
	})

	// test error name_book tidak boleh kurang dari 3 huruf
	testTable = append(testTable, testCase{
		name:      "validation name_book less than 5 letters",
		wantError: true,
		idBook:    1,
		input: entity.Book{
			NameBook: "Yo",
			Author:   "Inayatulloh",
		},
		expectedError: errors.New("name_book_less_than_5_letters"),
	})

	// test error author tidak boleh kurang dari 3 huruf
	testTable = append(testTable, testCase{
		name:      "validation author less than 5 letters",
		wantError: true,
		idBook:    1,
		input: entity.Book{
			NameBook: "Muiin's Book",
			Author:   "Da",
		},
		expectedError: errors.New("author_name_less_than_5_letters"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Buat mock controller
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepository(mockController)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			// Ambil hasil
			res, err := service.UpdateBook(testCase.idBook, testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestBookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult error
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepository)
	}

	var testTable []testCase

	// test success deleted
	testTable = append(testTable, testCase{
		name:           "success",
		wantError:      false,
		expectedResult: nil,
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{}, nil).Times(1)
			mock.EXPECT().DeleteOne(gomock.Any()).Return(nil).Times(1)
		},
	})

	// test error book no exist
	testTable = append(testTable, testCase{
		name:          "book not exist",
		wantError:     true,
		expectedError: errors.New("not_found"),
		onBookRepo: func(mock *mocks.MockBookRepository) {
			mock.EXPECT().GetOne(gomock.Any()).Return(entity.Book{}, gorm.ErrRecordNotFound)
			mock.EXPECT().DeleteOne(gomock.Any()).Return(gorm.ErrRecordNotFound).AnyTimes()
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			bookRepo := mocks.NewMockBookRepository(mockController)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
