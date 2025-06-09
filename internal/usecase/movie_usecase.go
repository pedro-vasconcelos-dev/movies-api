package usecase

import (
	"fmt"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
)

type MovieUsecase struct {
	Repo domain.MovieRepository
}

func (u *MovieUsecase) GetAllMovies() ([]domain.Movie, error) {
	return u.Repo.GetAll()
}

func (u *MovieUsecase) GetMovieByID(id int) (*domain.Movie, error) {
	return u.Repo.GetByID(id)
}

func (u *MovieUsecase) CreateMovie(movie *domain.Movie) error {
	existing, err := u.Repo.GetByTitle(movie.Title)
	if err == nil && existing != nil {
		return fmt.Errorf("movie title already exists")
	}
	if err != nil && err.Error() != "movie not found" {
		return err
	}
	return u.Repo.Create(movie)
}

func (u *MovieUsecase) UpdateMovie(id int, movie *domain.Movie) error {
	return u.Repo.Update(id, movie)
}

func (u *MovieUsecase) DeleteMovie(id int) error {
	return u.Repo.Delete(id)
}
