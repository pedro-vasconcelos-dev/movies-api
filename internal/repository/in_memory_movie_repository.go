package repository

import (
	"errors"

	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
)

type InMemoryMovieRepository struct {
	data map[int]domain.Movie
}

func NewInMemoryMovieRepository() *InMemoryMovieRepository {
	return &InMemoryMovieRepository{data: make(map[int]domain.Movie)}
}

func (r *InMemoryMovieRepository) GetAll() ([]domain.Movie, error) {
	movies := make([]domain.Movie, 0)
	for _, movie := range r.data {
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *InMemoryMovieRepository) GetByID(id int) (*domain.Movie, error) {
	movie, exists := r.data[id]
	if !exists {
		return nil, errors.New("movie not found")
	}
	return &movie, nil
}

func (r *InMemoryMovieRepository) Create(movie *domain.Movie) error {
	if _, exists := r.data[movie.ID]; exists {
		return errors.New("movie already exists")
	}
	r.data[movie.ID] = *movie
	return nil
}

func (r *InMemoryMovieRepository) Update(id int, movie *domain.Movie) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("movie not found")
	}
	// preserve the original ID so it doesn't get overwritten if the body
	// omits or changes the ID field
	movie.ID = id
	r.data[id] = *movie
	return nil
}

func (r *InMemoryMovieRepository) Delete(id int) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("movie not found")
	}
	delete(r.data, id)
	return nil
}
