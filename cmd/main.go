package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/pedro-vasconcelos-dev/movies-api/internal/delivery/gin"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/repository"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/usecase"
)

func main() {
	r := gin.Default()

	repo, err := repository.NewGORMMovieRepository(
		"localhost",
		5432,
		"postgres",
		"postgres",
		"movies",
	)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	usecase := usecase.MovieUsecase{Repo: repo}

	println("teste")
	routes.NewMovieHandler(r, &usecase)

	r.Run("localhost:8080")
}
