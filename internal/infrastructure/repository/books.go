package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Redarcher9/public-library/internal/domain"
	"github.com/Redarcher9/public-library/internal/infrastructure/models/tables"
	"gorm.io/gorm"
)

type Books struct {
	gormDB *gorm.DB
}

func NewBooksRepo(gormDB *gorm.DB) *Books {
	return &Books{
		gormDB: gormDB,
	}
}

func (b *Books) GetBooks(ctx context.Context, offset, limit int) ([]*domain.Book, error) {

	var books []*tables.Books
	result := b.gormDB.Model(&tables.Books{}).
		Limit(limit).
		Offset(offset).
		Find(&books)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// if book not found return nil
			return nil, nil
		}
		return nil, result.Error
	}
	if len(books) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	domainBooks := make([]*domain.Book, 0, len(books))
	for _, b := range books {
		domainBooks = append(domainBooks, b.ToDomain())
	}

	return domainBooks, nil
}

func (b *Books) GetBookByID(ctx context.Context, ID int) (*domain.Book, error) {

	var book tables.Books // Note: Not a pointer here
	result := b.gormDB.
		Where("id = ?", ID).
		First(&book)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound // Return nil if not found
		}
		return nil, fmt.Errorf("failed to get book by ID: %w", result.Error)
	}

	return book.ToDomain(), nil
}

func (b *Books) CreateBook(ctx context.Context, book *domain.Book) error {
	// Check if the book already exists (by Title and Author)
	var existingBook tables.Books
	if err := b.gormDB.Where("title = ? AND author = ?", book.Title, book.Author).First(&existingBook).Error; err == nil {
		return gorm.ErrDuplicatedKey
	}

	newBook := &tables.Books{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	if err := b.gormDB.Create(newBook).Error; err != nil {
		return err
	}

	return nil
}

func (b *Books) UpdateBookByID(ctx context.Context, ID int, book domain.Book) error {
	response := b.gormDB.Model(&tables.Books{}).Where("id = ?", ID).Updates(book)
	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (b *Books) DeleteBookByID(ctx context.Context, ID int) error {
	response := b.gormDB.Where("id = ?", ID).Delete(&tables.Books{})
	if response.Error != nil {
		return response.Error
	}
	return nil
}
