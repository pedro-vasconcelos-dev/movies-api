package domain

type MovieRepository interface {
	GetAll() ([]Movie, error)
	GetByID(id int) (*Movie, error)
	GetByTitle(title string) (*Movie, error)
	Create(movie *Movie) error
	Update(id int, movie *Movie) error
	Delete(id int) error
}
