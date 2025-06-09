package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/pedro-vasconcelos-dev/movies-api/internal/delivery/gin"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/repository"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/usecase"
)

func main() {
	r := gin.Default()

	repo, err := repository.NewPostgresMovieRepository(
		"54.207.134.101",
		6969,
		"gold_pedro",
		"@0r53p_dlog%",
		"RPA_Zap",
	)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	usecase := usecase.MovieUsecase{Repo: repo}

	routes.NewMovieHandler(r, &usecase)

	r.Run("localhost:8080")
}
