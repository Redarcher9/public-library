package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Redarcher9/public-library/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	BookInteractor BookService
}

func NewBookController(bookService BookService) *BookController {
	if bookService == nil {
		return nil
	}
	return &BookController{
		BookInteractor: bookService,
	}
}

func (bc *BookController) GetBooks(g *gin.Context) {
	// Get Query parameters with default values
	offset, err := strconv.Atoi(g.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(g.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	books, err := bc.BookInteractor.GetBooks(g, offset, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
		})
		return
	}
	g.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(g *gin.Context) {
	// Get the 'ID' parameter
	IDParam := g.Param("id")

	// Convert 'ID' to Int
	id, err := strconv.Atoi(IDParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid ID format",
		})
		return
	}

	book, err := bc.BookInteractor.GetBookByID(g, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: fmt.Sprintf("Book for ID %d not found", id),
			})
			return
		}
		g.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
		})
		return
	}
	g.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBookByID(g *gin.Context) {
	// Get the 'ID' parameter
	IDParam := g.Param("id")

	// Convert 'ID' to Int
	id, err := strconv.Atoi(IDParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid ID format",
		})
		return
	}

	err = bc.BookInteractor.DeleteBookByID(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
		})
		return
	}
	g.Status(http.StatusOK)
}

func (bc *BookController) UpdateBookByID(g *gin.Context) {
	// Get the 'ID' parameter
	IDParam := g.Param("id")

	// Convert 'ID' to Int
	id, err := strconv.Atoi(IDParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid ID format",
		})
		return
	}

	var req domain.Book
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: fmt.Sprintf("Invalid data: %s", err.Error()),
		})
		return
	}

	// Validate the input data using the Book's Validate method
	if err := req.Validate(); err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: fmt.Sprintf("Invalid data: %s", err.Error()),
		})
		return
	}

	// call Service for updating
	err = bc.BookInteractor.UpdateBookByID(g, id, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			g.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: fmt.Sprintf("Book for ID %d not found", id),
			})
			return
		}
		g.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "book updated successfully"})
}

func (bc *BookController) CreateBook(g *gin.Context) {
	var req domain.Book
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: fmt.Sprintf("Invalid data: %s", err.Error()),
		})
		return
	}

	// Validate the input data using the Book's Validate method
	if err := req.Validate(); err != nil {
		g.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: fmt.Sprintf("Invalid data: %s", err.Error()),
		})
		return
	}

	err := bc.BookInteractor.CreateBook(g, &req)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			g.JSON(http.StatusConflict, domain.ErrorResponse{
				Message: fmt.Sprintf("Book with Title %s and Author %s already exists", req.Title, req.Author),
			})
			return
		}
		g.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, req)
}
