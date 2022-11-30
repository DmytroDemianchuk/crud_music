package domain

import (
	"errors"
	"time"
)

var (
	ErrBookNotFound = errors.New("book not found")
)

type Music struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Performer   string    `json:"performer"`
	RealiseDate time.Time `json:"realise_date"`
	Genre       string    `json:"genre"`
}

type UpdateBookInput struct {
	Name        *string    `json:"name"`
	Performer   *string    `json:"performer"`
	RealiseDate *time.Time `json:"realise_date"`
	Genre       *string    `json:"genre"`
}
