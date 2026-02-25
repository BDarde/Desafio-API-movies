package main

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

var mvs []Movie
var errMinVote = errors.New("filtro de votos minimos nao poder ser maior que máximo")

// busca de N filmes com filtro
func (app *App) HandlerFindMovies(c *gin.Context) {

	filter, err := extractFilter(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Message: "Não foi possível reconhecer as informações do filtro de busca",
			Error:   err.Error(),
		})
		return
	}

	fmt.Printf("filtro extraído %+v", filter)

	mvsResonse := []MovieResponse{}

	for _, mv := range app.Movies {

		if filter.Genre != "" && mv.Genre != filter.Genre {
			continue
		}

		year := strings.Split(mv.ReleaseDate, "-")[0]
		if filter.ReleaseYear != "" && filter.ReleaseYear != year {
			continue
		}

		// alteração na logica
		if filter.Min != 0 && mv.VoteAverage < filter.Min {
			continue
		}

		if filter.Max != 0 && mv.VoteAverage > filter.Max {
			continue
		}

		if filter.Min > filter.Max {
			c.JSON(http.StatusBadRequest, ResponseError{
				Message: " Conflito nos filtros de votos",
				Error:   errMinVote.Error(),
			})
			return
		}
		//

		mvsResonse = append(mvsResonse, MovieResponse{ID: mv.ID,
			Title:       mv.Title,
			ReleaseDate: mv.ReleaseDate,
			Genre:       mv.Genre,
			VoteAverage: mv.VoteAverage})

	}

	mvsResonse = paginate(mvsResonse, filter.Page, filter.Limit)

	c.JSON(http.StatusOK, mvsResonse)

}

func (app *App) HandlerFindMovie(c *gin.Context) {

	var movie MovieWithMetrics
	var profit int64

	// busca individual
	fmt.Printf("handler de busca individual")
	id, prs := c.Params.Get("id")
	if prs {

		idInt, _ := strconv.Atoi(id)
		for _, mv := range app.Movies {
			if mv.ID == idInt {
				profit = mv.Revenue - mv.Budget
				movie = MovieWithMetrics{
					Movie: mv,
					Metrics: Metrics{
						Profit: profit,
						ROI:    float64(profit) / float64(mv.Budget),
						// score = (vote_average * 0.6) + (lucro * 0.4)
						SuccessScore: (movie.VoteAverage * 0.6) + (float64(profit) * 0.4),
					},
				}

				c.JSON(http.StatusOK, movie)
			}
		}
	}

}

func (app *App) HandlerAnalycts(c *gin.Context) {

	response, err := calc(app.Movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Message: "Não foi possível terminar os calculos",
			Error:   err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)

}

func (app *App) HandlerTopStudios(c *gin.Context) {

	response := calcStudios(app.Movies)

	sort.Slice(response, func(i, j int) bool {
		return response[i].TotalProfit > response[j].TotalProfit
	})

	for i := range response {
		response[i].Ranking = i + 1
	}

	c.JSON(http.StatusOK, response)
}

func (app *App) HandlerGenreStatus(c *gin.Context) {

	response := calcGenre(app.Movies)

	c.JSON(http.StatusOK, response)
}
