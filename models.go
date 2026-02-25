package main

// Atributos de cada filme r
type Movie struct {
	Budget            int64   `json:"budget"`
	Homepage          string  `json:"homepage"`
	ID                int     `json:"id"`
	OriginalLanguage  string  `json:"original_language"`
	Overview          string  `json:"overview"`
	Popularity        float64 `json:"popularity"`
	ReleaseDate       string  `json:"release_date"`
	Revenue           int64   `json:"revenue"`
	Runtime           int     `json:"runtime"`
	Title             string  `json:"title"`
	VoteAverage       float64 `json:"vote_average"`
	VoteCount         int     `json:"vote_count"`
	Genre             string  `json:"genre"`
	ProductionCompany string  `json:"production_company"`
}

// Metricas (indicadores individuais)
type Metrics struct {
	Profit       int64   `json:"profit"`
	ROI          float64 `json:"roi"`
	SuccessScore float64 `json:"success_score"`
}
type MovieWithMetrics struct {
	Movie
	Metrics
}

// Filtro de Pesquisa

type Filter struct {
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	Genre       string  `json:"genre"`
	ReleaseYear string  `json:"release_year"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
}

// Estrutura de resposta da consulta

type Metadata struct {
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
}

type MovieResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Genre       string  `json:"genre"`
	VoteAverage float64 `json:"vote_average"`
}

type APIResponse struct {
	Metadata Metadata        `json:"metadata"`
	Data     []MovieResponse `json:"data"`
}

// Struct do calculo geral da base
type Summary struct {
	TotalMoviesProcessed int     `json:"total_movies_processed"`
	GlobalAverageRating  float64 `json:"global_average_rating"`
	TotalRevenue         int64   `json:"total_revenue"`
	TotalBudget          int64   `json:"total_budget"`
	AverageRevenue       int64   `json:"average_revenue"`
	AverageBudget        int64   `json:"average_budget"`
	AverageProfit        int64   `json:"average_profit"`
	OverallROI           float64 `json:"overall_roi"`
}

type TopPerformingYear struct {
	Year         int   `json:"year"`
	TotalRevenue int64 `json:"total_revenue"`
}

type ResponseAnaytics struct {
	Summary           Summary
	TopPerformingYear TopPerformingYear
}

// struct de calculo de estudios

type StudioStats struct {
	Ranking     int     `json:"ranking"`
	Studio      string  `json:"studio"`
	AverageROI  float64 `json:"average_roi"`
	TotalProfit float64 `json:"total_profit"`
	MovieCount  int     `json:"movie_count"`
}

type GenreStats struct {
	Genre              string  `json:"genre"`
	AverageVote        float64 `json:"average_vote"`
	AverageROI         float64 `json:"average_roi"`
	MovieCount         int     `json:"movie_count"`
	TotalRevenue       int64   `json:"total_revenue"`
	MarketSharePercent float64 `json:"market_share_percent"`
}
