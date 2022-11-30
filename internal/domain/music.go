package domain

import (
	"errors"
	"time"
)

var (
	ErrMusicNotFound = errors.New("music not found")
)

type Music struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Performer   string    `json:"performer"`
	RealiseDate time.Time `json:"realise_date"`
	Genre       string    `json:"genre"`
}

type UpdateMusicInput struct {
	Name        *string    `json:"name"`
	Performer   *string    `json:"performer"`
	RealiseDate *time.Time `json:"realise_date"`
	Genre       *string    `json:"genre"`
}
