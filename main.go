package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ainmtsn1999/orm-book-api-test/config"
	"github.com/ainmtsn1999/orm-book-api-test/model/web"
	"github.com/ainmtsn1999/orm-book-api-test/repository"
	"github.com/ainmtsn1999/orm-book-api-test/routes"
	"github.com/ainmtsn1999/orm-book-api-test/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := config.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(config.NewGorm.DB)
	serv := service.NewService(repo)

	newRouter := gin.New()
	routes.BookRouter(newRouter, serv)
	newRouter.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, web.BookResponse{Message: "Page not found"})
	})

	port := os.Getenv("PORT")
	err = newRouter.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
