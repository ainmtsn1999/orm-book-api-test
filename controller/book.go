package controller

import (
	"github.com/ainmtsn1999/orm-book-api-test/helper"
	"github.com/ainmtsn1999/orm-book-api-test/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController interface {
	CreateBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

func (c *Controller) CreateBook(ctx *gin.Context) {
	var inputBook entity.Book
	err := ctx.ShouldBindJSON(&inputBook)
	if err != nil {
		helper.ResponseError(ctx, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = inputBook.Validation()
	if err != nil {
		helper.ValidationError(ctx, err)
		return
	}
	book, err := c.service.CreateBook(inputBook)
	if err != nil {
		helper.CustomResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func (c *Controller) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		helper.CustomResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *Controller) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}
	book, err := c.service.GetBook(cvtId)
	if err != nil {
		helper.CustomResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c *Controller) UpdateBook(ctx *gin.Context) {
	// Get id from param
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	// Get update from json
	var inputBook entity.Book
	err = ctx.ShouldBindJSON(&inputBook)
	if err != nil {
		helper.ResponseError(ctx, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = inputBook.Validation()
	if err != nil {
		helper.ValidationError(ctx, err)
		return
	}

	book, err := c.service.UpdateBook(cvtId, inputBook)
	if err != nil {
		helper.CustomResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *Controller) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	cvtId, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(ctx, "ID params not valid", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteBook(cvtId)
	if err != nil {
		helper.CustomResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
