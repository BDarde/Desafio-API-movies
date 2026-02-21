package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	path string
	mvs  []Movie
)

const NAME_FILE = "movies"

func main() {

	//le do arquivo, se estiver vazio preenche
	_, err := os.ReadFile(NAME_FILE)
	if err == io.EOF {
		// arquivo vazio, é necessário alimentar o disco
		//extração dos dados do arquivo via argumento de linha de comando
		if len(os.Args) > 0 {
			path = os.Args[1]
		} else {
			panic(fmt.Errorf("é necessário o caminho do arquivo"))
		}
		// abre o arquivo vindo de path
		f, err := os.Open(path)
		if err != nil {
			panic(fmt.Errorf("não foi possível ler o arquivo %s", err.Error()))
		}
		// passa os dados para um slice da struct dos dados, conforme pré configurada em config
		mvs, err := importData(f)
		if err != nil {
			panic(err)
		}
		//perssiste os dados no disco
		persistJson(mvs)
	}

	// sincroniza os filmes com a struct
	update(NAME_FILE, &mvs)

	initServe()

}

func initServe() {
	e := gin.Default()
	e.GET("/movies", HandlerFindMovies)
	e.GET("/movies/:id", HandlerFindMovie)
	e.GET("/analytics/dashboard", HandlerAnalycts)
	e.GET("/analytics/top-studios", HandlerTopStudios)
	e.GET("/analytics/genre-stats", HandlerGenreStatus)
	e.Run(":8080")
}

////
