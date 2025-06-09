package repository

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
)

// GORMMovieRepository implements MovieRepository using GORM for SQL Server.
type GORMMovieRepository struct {
	db *gorm.DB
}

// NewGORMMovieRepository initializes GORM with SQL Server dialect.
func NewGORMMovieRepository(host string, port int, user, password, dbname string) (*GORMMovieRepository, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, host, port, dbname)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &GORMMovieRepository{db: db}, nil
}

func (r *GORMMovieRepository) GetAll() ([]domain.Movie, error) {
	var movies []domain.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *GORMMovieRepository) GetByID(id int) (*domain.Movie, error) {
	var movie domain.Movie
	if err := r.db.First(&movie, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, err
	}
	return &movie, nil
}

func (r *GORMMovieRepository) GetByTitle(title string) (*domain.Movie, error) {
	var movie domain.Movie
	if err := r.db.Where("title = ?", title).First(&movie).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, err
	}
	return &movie, nil
}

func (r *GORMMovieRepository) Create(movie *domain.Movie) error {
	return r.db.Create(movie).Error
}

func (r *GORMMovieRepository) Update(id int, movie *domain.Movie) error {
	movie.ID = id
	res := r.db.Model(&domain.Movie{}).Where("id = ?", id).Updates(movie)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("movie not found")
	}
	return nil
}

func (r *GORMMovieRepository) Delete(id int) error {
	res := r.db.Delete(&domain.Movie{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("movie not found")
	}
	return nil
}

func (r *GORMMovieRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
