package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const FILE_TESTE = "moviesJson.test"

func setupTestApp(t *testing.T) *App {
	t.Helper()

	app, err := falseMainTestIntegration()
	if err != nil {
		t.Fatalf("Erro no setup do app: %v", err)
	}

	// Agenda a remoção do arquivo de teste assim que o teste terminar
	t.Cleanup(func() {
		os.Remove(FILE_TESTE)
	})

	return app
}

func falseMainTestIntegration() (*App, error) {
	app := &App{}

	mvs, err := loadMovies("movies.csv")
	if err != nil {
		return nil, fmt.Errorf("falhou ao ler o arquivo: %w", err)
	}

	app.Movies = mvs

	err = createJsonFile(app.Movies, FILE_TESTE)
	if err != nil {
		return nil, fmt.Errorf("falhou ao persistir dados json: %w", err)
	}

	return app, nil
}

func TestGetMoviesIntegration(t *testing.T) {

	a, err := falseMainTestIntegration()
	if err != nil {
		t.Fatal(err)
	}
	router := setupRouter(a)

	req, _ := http.NewRequest("GET", "/movies", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperava o codigo de sucesso %d", w.Code)
	}

}

func TestGetCalcMoviesIntegration(t *testing.T) {

	response := ResponseAnaytics{}

	app, err := falseMainTestIntegration()
	if err != nil {
		t.Fatal(err)
	}
	router := setupRouter(app)

	req, _ := http.NewRequest("GET", "/analytics/dashboard", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperava %d mas recebeu %d. Body: %s",
			http.StatusOK,
			w.Code,
			w.Body.String(),
		)
	}
	if len(w.Body.String()) == 0 {
		t.Fatal("esperava conteúdo no body da requisição")
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("não conseguiu fazer o unmarshal dos dados")

	}

	if response.Summary.TotalRevenue < 4 {
		t.Fatal("esperava numero coerente")
	}

}
