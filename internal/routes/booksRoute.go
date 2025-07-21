package routes

import (
	"github.com/Redarcher9/public-library/internal/controller"
	"github.com/Redarcher9/public-library/internal/infrastructure/repository"
	"github.com/Redarcher9/public-library/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewBookRouter(group *gin.RouterGroup, db *gorm.DB) {
	//Instantiate Repository, Service and Controller through dependency injection

	bookRepo := repository.NewBooksRepo(db)
	bookService := service.NewBookInteractor(bookRepo)
	bookController := controller.NewBookController(bookService)

	//Initialise Routes
	group.GET("/ping")
	group.GET("/books", bookController.GetBooks)
	group.GET("/books/:id", bookController.GetBookByID)
	group.DELETE("/books/:id", bookController.DeleteBookByID)
	group.PUT("/books/:id", bookController.UpdateBookByID)
	group.POST("/books", bookController.CreateBook)

}
