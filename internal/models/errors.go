package models

import "errors"

var (
	ErrNotFound     = errors.New("Product not found")
	ErrInvalidInput = errors.New("Invalid input")
	ErrDatabase     = errors.New("Database error")
	ErrStorageFull  = errors.New("storage limit reached")
)
