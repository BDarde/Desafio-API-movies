package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// nome do arquivo em disco
const NAME_FILE = "movies"

type App struct {
	Movies []Movie
}

var (
	state App
	path  string
)

func main() {
	// Tenta ler o arquivo JSON existente
	data, err := os.ReadFile(NAME_FILE)

	// se o arquivo não existe ou está vazio
	if err != nil || len(data) == 0 {
		fmt.Println("JSON não encontrado ou vazio. Carregando do CSV bruto")

		// verifica argumentos
		if len(os.Args) < 2 {
			panic("é necessário passar o caminho do arquivo CSV: go run . movies.csv")
		}

		path := os.Args[1]
		fmt.Printf("nome do arquivo de input capturado: %s", path)

		// carrega do CSV
		state.Movies, err = loadMovies(path)
		if err != nil {
			panic(fmt.Errorf("falhou ao ler o arquivo: %w", err))
		}

		// persiste no disco pela primeira vez
		err = persistJson(state.Movies, NAME_FILE)
		if err != nil {
			panic(fmt.Errorf("falhou ao persistir JSON: %w", err))
		}
	} else {
		// se o arquivo existe e tem dados, faz o Unmarshal para a struct
		err = json.Unmarshal(data, &state.Movies)
		if err != nil {
			panic(fmt.Errorf("falhou ao sincronizar JSON: %w", err))
		}
		fmt.Println("Dados carregados do cache JSON com sucesso.")
	}

	e := setupRouter(&state)
	e.Run(":8080")
}

func setupRouter(a *App) *gin.Engine {
	e := gin.Default()
	e.GET("/movies", a.HandlerFindMovies)
	e.GET("/movies/:id", a.HandlerFindMovie)
	e.GET("/analytics/dashboard", a.HandlerAnalycts)
	e.GET("/analytics/top-studios", a.HandlerTopStudios)
	e.GET("/analytics/genre-stats", a.HandlerGenreStatus)
	return e
}
