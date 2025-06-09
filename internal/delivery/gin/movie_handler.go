package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/domain"
	"github.com/pedro-vasconcelos-dev/movies-api/internal/usecase"
)

func NewMovieHandler(r *gin.Engine, usecase *usecase.MovieUsecase) {
	r.GET("/movies", func(c *gin.Context) {
		movies, err := usecase.GetAllMovies()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, movies)
	})

	r.GET("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		movie, err := usecase.GetMovieByID(id)
		if err != nil {
			if err.Error() == "movie not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, movie)
	})

	r.POST("/movies", func(c *gin.Context) {
		var movie domain.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := usecase.CreateMovie(&movie); err != nil {
			if err.Error() == "movie title already exists" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusCreated, movie)
	})

	r.PUT("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedMovie domain.Movie
		if err := c.ShouldBindJSON(&updatedMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := usecase.UpdateMovie(id, &updatedMovie); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Filme não encontrado
			return
		}

		c.JSON(http.StatusOK, updatedMovie)
	})

	r.DELETE("/movies/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := usecase.DeleteMovie(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})
}
