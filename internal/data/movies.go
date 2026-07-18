package data

import (
	"time"

	"github.com/shodruzhoshimzoda/Greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // (hiphen will make disappear this field)
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   []string  `json:"version,omitempty"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than or equal to 2018")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must be greater than or equal to current year")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must not be positive number")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) > 1, "genres", "genres must contain at least 1 item")
	v.Check(len(movie.Genres) < 5, "genres", "genres must not contain at least 5 item")
	v.Check(validator.Unique(movie.Genres), "genres", "genres must not contain duplicate values")
}
