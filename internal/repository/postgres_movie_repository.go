package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
)

type PostgresMovieRepository struct {
	db *sql.DB
}

func NewPostgresMovieRepository(host string, port int, user, password, dbname string) (*PostgresMovieRepository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresMovieRepository{db: db}, nil
}

func (r *PostgresMovieRepository) GetAll() ([]domain.Movie, error) {
	rows, err := r.db.Query("SELECT id, title, director, genre, year FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		var m domain.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Director, &m.Genre, &m.Year); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *PostgresMovieRepository) GetByID(id int) (*domain.Movie, error) {
	var m domain.Movie
	row := r.db.QueryRow("SELECT id, title, director, genre, year FROM movies WHERE id=$1", id)
	if err := row.Scan(&m.ID, &m.Title, &m.Director, &m.Genre, &m.Year); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, err
	}
	return &m, nil
}

func (r *PostgresMovieRepository) Create(movie *domain.Movie) error {
	_, err := r.db.Exec("INSERT INTO movies (id, title, director, genre, year) VALUES ($1, $2, $3, $4, $5)", movie.ID, movie.Title, movie.Director, movie.Genre, movie.Year)
	return err
}

func (r *PostgresMovieRepository) Update(id int, movie *domain.Movie) error {
	res, err := r.db.Exec("UPDATE movies SET title=$1, director=$2, genre=$3, year=$4 WHERE id=$5", movie.Title, movie.Director, movie.Genre, movie.Year, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("movie not found")
	}
	movie.ID = id
	return nil
}

func (r *PostgresMovieRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM movies WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("movie not found")
	}
	return nil
}

func (r *PostgresMovieRepository) Close() error {
	return r.db.Close()
}
