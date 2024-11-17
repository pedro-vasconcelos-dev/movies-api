package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/pedro-vasconcelos-dev/movies-api/internal/delivery/gin"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/repository"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/usecase"
)

func main() {
	r := gin.Default()

	repo := repository.NewInMemoryMovieRepository()
	usecase := usecase.MovieUsecase{Repo: repo}

	routes.NewMovieHandler(r, &usecase)

	r.Run("localhost:8080")
}
