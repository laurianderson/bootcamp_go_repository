package store

import (
	"errors"
	"github.com/laurianderson/bootcamp_go_repository/internal/domain"
)

// Errrors
var (
	ErrRepositoryNotFound   = errors.New("product not found")
	ErrRepositoryInternal   = errors.New("an internal error")
	ErrRepositoryDuplicated = errors.New("duplicated product")
)

type StoreInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(product domain.Product) error
	// Update actualiza un producto
	Update(product domain.Product) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
