package helper

import (
	"github.com/ainmtsn1999/orm-book-api-test/model/web"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

func CustomResponseError(ctx *gin.Context, err error) {
	switch err.Error() {
	case "already_exist":
		ResponseError(ctx, "This book already exists", http.StatusBadRequest)
		break
	case "no_data":
		ResponseError(ctx, "No books", http.StatusOK)
		break
	case "not_found":
		ResponseError(ctx, "Book not found", http.StatusNotFound)
		break
	default:
		ResponseError(ctx, "Server Error", http.StatusInternalServerError)
		break
	}
}

func ResponseError(ctx *gin.Context, msg any, code int) {
	ctx.AbortWithStatusJSON(code, web.BookResponse{
		Message: msg,
	})
}

func ValidationError(ctx *gin.Context, err error) {
	var errorSlice []gin.H
	if e, ok := err.(validation.Errors); ok {
		for field, msg := range e {
			errorSlice = append(errorSlice, gin.H{
				field: msg.Error(),
			})
		}
	}
	ctx.AbortWithStatusJSON(400, gin.H{
		"validation": errorSlice,
	})
}
