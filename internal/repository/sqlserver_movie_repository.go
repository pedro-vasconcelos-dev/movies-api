package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
)

// SQLServerMovieRepository implements the MovieRepository interface using
// Microsoft SQL Server as the backend.
type SQLServerMovieRepository struct {
	db *sql.DB
}

// NewSQLServerMovieRepository creates a new repository connected to a SQL Server database.
func NewSQLServerMovieRepository(host string, port int, user, password, dbname string) (*SQLServerMovieRepository, error) {
	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s", host, port, user, password, dbname)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &SQLServerMovieRepository{db: db}, nil
}

func (r *SQLServerMovieRepository) GetAll() ([]domain.Movie, error) {
	rows, err := r.db.Query("SELECT id, title, director, genre, year FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []domain.Movie{}
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

func (r *SQLServerMovieRepository) GetByID(id int) (*domain.Movie, error) {
	var m domain.Movie
	row := r.db.QueryRow("SELECT id, title, director, genre, year FROM movies WHERE id = @p1", id)
	if err := row.Scan(&m.ID, &m.Title, &m.Director, &m.Genre, &m.Year); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, err
	}
	return &m, nil
}

func (r *SQLServerMovieRepository) GetByTitle(title string) (*domain.Movie, error) {
	var m domain.Movie
	row := r.db.QueryRow("SELECT id, title, director, genre, year FROM movies WHERE title = @p1", title)
	if err := row.Scan(&m.ID, &m.Title, &m.Director, &m.Genre, &m.Year); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, err
	}
	return &m, nil
}

func (r *SQLServerMovieRepository) Create(movie *domain.Movie) error {
	return r.db.QueryRow(
		"INSERT INTO movies (title, director, genre, year) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3, @p4)",
		movie.Title,
		movie.Director,
		movie.Genre,
		movie.Year,
	).Scan(&movie.ID)
}

func (r *SQLServerMovieRepository) Update(id int, movie *domain.Movie) error {
	res, err := r.db.Exec("UPDATE movies SET title=@p1, director=@p2, genre=@p3, year=@p4 WHERE id=@p5",
		movie.Title, movie.Director, movie.Genre, movie.Year, id)
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

func (r *SQLServerMovieRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM movies WHERE id=@p1", id)
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

func (r *SQLServerMovieRepository) Close() error {
	return r.db.Close()
}
