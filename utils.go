package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// lista de constantes para verificação
const (
	page         = "page"
	limit        = "limit"
	genre        = "genre"
	release_year = "release_year"
	min_vote     = "min_vote"
	max_vote     = "max_vote"
)

// funcao responsável por preencher a struct com os dados do csv
func importData(f *os.File) ([]Movie, error) {

	var mvs []Movie

	//inicia o processo de leitura do arquivo csv
	read := csv.NewReader(f)

	_, err := read.Read()
	if err != nil {
		return nil, fmt.Errorf("arquivo sem header ou vazio %s", err.Error())
	}

	for {

		line, err := read.Read()
		if err == io.EOF {
			//fim do arquivo
			break
		}

		mv, err := parseMovie(line)
		if err != nil {
			continue
		}

		mvs = append(mvs, mv)

	}

	return mvs, nil

}

func parseMovie(line []string) (Movie, error) {

	if len(line) < 15 {
		return Movie{}, fmt.Errorf("linha inválida: colunas insuficientes")
	}

	budget, err := strconv.ParseInt(line[0], 10, 64)
	if err != nil {
		return Movie{}, fmt.Errorf("erro budget: %w", err)
	}

	id, err := strconv.Atoi(line[2])
	if err != nil {
		return Movie{}, fmt.Errorf("erro id: %w", err)
	}

	popularity, err := strconv.ParseFloat(line[5], 64)
	if err != nil {
		return Movie{}, fmt.Errorf("erro popularity: %w", err)
	}

	revenue, err := strconv.ParseInt(line[7], 10, 64)
	if err != nil {
		return Movie{}, fmt.Errorf("erro revenue: %w", err)
	}

	runtime, err := strconv.Atoi(line[8])
	if err != nil {
		return Movie{}, fmt.Errorf("erro runtime: %w", err)
	}

	voteAverage, err := strconv.ParseFloat(line[10], 64)
	if err != nil {
		return Movie{}, fmt.Errorf("erro vote_average: %w", err)
	}

	voteCount, err := strconv.Atoi(line[11])
	if err != nil {
		return Movie{}, fmt.Errorf("erro vote_count: %w", err)
	}

	return Movie{
		Budget:            budget,
		Homepage:          line[1],
		ID:                id,
		OriginalLanguage:  line[3],
		Overview:          line[4],
		Popularity:        popularity,
		ReleaseDate:       line[6],
		Revenue:           revenue,
		Runtime:           runtime,
		Title:             line[9],
		VoteAverage:       voteAverage,
		VoteCount:         voteCount,
		Genre:             line[12],
		ProductionCompany: line[13],
	}, nil
}

func paginate(items []MovieResponse, page, limit int) []MovieResponse {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(items) {
		return []MovieResponse{}
	}

	if end > len(items) {
		end = len(items)
	}

	return items[start:end]
}
func extractFilter(c *gin.Context) (Filter, error) {
	var f Filter

	query := c.Request.URL.Query()

	f.Genre = query.Get(genre)
	f.ReleaseYear = query.Get("release_year")

	if v := query.Get(limit); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return Filter{}, fmt.Errorf("limit inválido")
		}
		f.Limit = limit
	} else {
		f.Limit = 10
	}

	if v := query.Get(page); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return Filter{}, fmt.Errorf("page inválido")
		}
		f.Page = page
	} else {
		f.Page = 1
	}

	if v := query.Get(min_vote); v != "" {
		min, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Filter{}, fmt.Errorf("min_vote inválido")
		}
		f.Min = min
	}

	if v := query.Get(max_vote); v != "" {
		max, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Filter{}, fmt.Errorf("max_vote inválido")
		}
		f.Max = max
	}

	if f.Min != 0 && f.Max == 0 {
		f.Max = 10
	}

	return f, nil
}

func calc(mvs []Movie) (ResponseAnaytics, error) {
	var accBudget, accRevenue, accProfit, ROI int64
	var accVotes float64
	totalMovies := len(mvs)

	for _, mv := range mvs {
		accVotes += mv.VoteAverage
		accRevenue += mv.Revenue
		accBudget += mv.Budget
		accProfit += mv.Revenue - mv.Budget

		if mv.Budget == 0 {
			continue
		}
		ROI += (mv.Revenue - mv.Budget) / mv.Budget
	}

	topPerfoming, err := calcTopYear(mvs)
	if err != nil {
		return ResponseAnaytics{}, err
	}

	return ResponseAnaytics{
		Summary: Summary{
			TotalMoviesProcessed: totalMovies,
			GlobalAverageRating:  accVotes / float64(totalMovies),
			TotalRevenue:         accRevenue,
			TotalBudget:          accBudget,
			AverageRevenue:       accRevenue / int64(totalMovies),
			AverageBudget:        accBudget / int64(totalMovies),
			AverageProfit:        accProfit / int64(totalMovies),
			OverallROI:           float64(ROI),
		},
		TopPerformingYear: TopPerformingYear{
			Year:         topPerfoming.Year,
			TotalRevenue: topPerfoming.TotalRevenue,
		},
	}, nil

}

func calcTopYear(mvs []Movie) (TopPerformingYear, error) {

	var (
		m          = make(map[string]int64)
		topRevenue int64
		topYear    string
		finalYear  int
		err        error
	)
	for _, mv := range mvs {

		if len(mv.ReleaseDate) < 4 {
			continue
		}

		year := mv.ReleaseDate[:4]

		revenueYear := mv.Revenue
		m[year] += revenueYear
	}

	for k, v := range m {

		if v > topRevenue {
			topRevenue = v
			topYear = k
		}

	}

	finalYear, err = strconv.Atoi(topYear)
	if err != nil {
		fmt.Printf("falha ao converter ano %s", err.Error())
		return TopPerformingYear{}, err
	}

	return TopPerformingYear{
		Year:         finalYear,
		TotalRevenue: topRevenue,
	}, nil

}
func calcStudios(mvs []Movie) []StudioStats {

	fmt.Println("entrou no calculo do studio")

	r := []StudioStats{}

	m := make(map[string]*StudioStats)

	fmt.Println("iterando sobre filmes")
	for _, mv := range mvs {

		if mv.ProductionCompany == "" {
			continue
		}

		if _, prs := m[mv.ProductionCompany]; !prs {
			m[mv.ProductionCompany] = &StudioStats{}
			fmt.Println("criando grupo do studio: ", mv.ProductionCompany)
		}

		acc := m[mv.ProductionCompany]

		profit := mv.Revenue - mv.Budget

		acc.TotalProfit += float64(profit)
		if mv.Budget != 0 {
			acc.AverageROI += float64(profit) / float64(mv.Budget)
		}
		acc.MovieCount++
	}

	for k, v := range m {

		averageROI := 0.0
		if v.MovieCount > 0 {
			averageROI = v.AverageROI / float64(v.MovieCount)
		}

		studio := StudioStats{
			Studio:      k,
			AverageROI:  averageROI,
			TotalProfit: v.TotalProfit,
			MovieCount:  v.MovieCount,
		}

		r = append(r, studio)
	}

	return r
}

func calcGenre(mvs []Movie) []GenreStats {

	var r []GenreStats

	m := make(map[string]*GenreStats)

	for _, mv := range mvs {

		if mv.Genre == "" {
			continue
		}
		if _, prs := m[mv.Genre]; !prs {
			m[mv.Genre] = &GenreStats{}
			fmt.Println("criando grupo de genero: ", mv.Genre)
		}

		acc := m[mv.Genre]
		//calculo final na atribuição da struct
		acc.TotalRevenue += mv.Revenue
		//calculo final na atribuição da struct
		acc.AverageVote += mv.VoteAverage

		profit := mv.Revenue - mv.Budget

		if mv.Budget != 0 {
			//calculo final na atribuição da struct
			acc.AverageROI += float64(profit) / float64(mv.Budget)
		}

		/// calculo final = divisao por quantidade de filmes do genero
		acc.MovieCount++

	}

	for k, v := range m {
		//montagem da struct por genero
		r = append(r, GenreStats{
			Genre:        k,
			AverageVote:  v.AverageVote / float64(v.MovieCount),
			AverageROI:   v.AverageROI / float64(v.MovieCount),
			MovieCount:   v.MovieCount,
			TotalRevenue: v.TotalRevenue,
			//(movie_count_do_genero / total_movies) * 100
			MarketSharePercent: (float64(v.MovieCount) / float64(len(mvs))) * 100,
		})

	}

	return r

}

func loadMovies(fileName string) ([]Movie, error) {

	var mvs []Movie

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fmt.Print("arquivo aberto com sucesso")

	mvs, err = importData(f)
	if err != nil {
		return nil, err
	}

	if len(mvs) == 0 {
		return nil, fmt.Errorf("não conseguiu processar os arquivo, o objeto json está vazio")
	}

	//** fechamento do arquivo economia de memoria
	defer f.Close()
	fmt.Printf("objeto json preenchido com sucesso")

	return mvs, nil
}

func persistJson(mvs []Movie, nameFile string) error {

	f, err := os.Create(nameFile)
	if err != nil {
		return err
	}

	b, err := json.Marshal(mvs)
	if err != nil {
		return err
	}

	n, err := f.Write(b)
	if err != nil || n == 0 {
		return err
	}

	return nil

}

func createJsonFile(mvs []Movie, nameFile string) error {
	f, err := os.Create(nameFile)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(mvs)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}
