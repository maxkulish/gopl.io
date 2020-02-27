package imdb

const (
	ApiURL = "http://www.omdbapi.com"
	ApiKey = "ccc457c6"
)

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
	Type   string `json:"Type"`
	ImdbID string `json:"imdbID"`
}

type SearchResult struct {
	TotalResults string   `json:"totalResults"`
	Movies       []*Movie `json:"Search"`
}
