package usecase

import (
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
	return u.Repo.Create(movie)
}

func (u *MovieUsecase) UpdateMovie(id int, movie *domain.Movie) error {
	return u.Repo.Update(id, movie)
}

func (u *MovieUsecase) DeleteMovie(id int) error {
	return u.Repo.Delete(id)
}
