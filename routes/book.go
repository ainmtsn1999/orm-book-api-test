package routes

import (
	"github.com/ainmtsn1999/orm-book-api-test/controller"
	"github.com/ainmtsn1999/orm-book-api-test/service"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, service service.ServiceInterface) {
	// Panggil controller
	handler := controller.NewController(service)
	api := router.Group("/books")
	{
		api.POST("", handler.CreateBook)       // Create book
		api.GET("", handler.GetAllBooks)       // Get all books
		api.GET("/:id", handler.GetBook)       // Get id book
		api.PUT("/:id", handler.UpdateBook)    // Update book
		api.DELETE("/:id", handler.DeleteBook) // Delete book
	}
}
