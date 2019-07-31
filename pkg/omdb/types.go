package omdb

type SearchResponse struct {
	Search       []MovieItem `json:"Search,omitempty"`
	TotalResults string      `json:"totalResults,omitempty"`
	Response     string      `json:"Response"`
	Error        string      `json:"Error,omitempty"`
}

type MovieItem struct {
	Title  string `json:"Title,omitempty" form:"title"`
	Year   string `json:"Year,omitempty" form:"year"`
	ImdbID string `json:"imdbID,omitempty" form:"imdbID"`
	Type   string `json:"Type,omitempty" form:"type"`
	Poster string `json:"Poster,omitempty" form:"poster"`
}

type MovieDetails struct {
	Title      string   `json:"Title,omitempty"`
	Year       string   `json:"Year,omitempty"`
	Rated      string   `json:"Rated,omitempty"`
	Released   string   `json:"Released,omitempty"`
	Runtime    string   `json:"Runtime,omitempty"`
	Genre      string   `json:"Genre,omitempty"`
	Director   string   `json:"Director,omitempty"`
	Writer     string   `json:"Writer,omtiempty"`
	Actors     string   `json:"Actors,omitempty`
	Plot       string   `json:"Plot,omitempty"`
	Language   string   `json:"Language,omitempty"`
	Country    string   `json:"Country,omitempty"`
	Awards     string   `json:"Awards,omitempty"`
	Poster     string   `json:"Poster,omitempty`
	Ratings    []Rating `json:"Ratings,omitempty`
	Metascore  string   `json:"Metascore,omitempty"`
	ImdbRating string   `json:"imdbRating,omitempty"`
	ImdbVotes  string   `json:"imdbVotes,omitempy"`
	ImdbID     string   `json:"imdbID,omitempty"`
	DVD        string   `json:"DVD,omitempy"`
	BoxOffice  string   `json:"BoxOffice,omitempty"`
	Production string   `json:"Production,omitempty"`
	Website    string   `json:"Website,omitempty"`
	Response   string   `json:"Response"`
	Error      string   `json:"Error,omitempty"`
}

type Rating struct {
	Source string `json:"Source,omitempty`
	Value  string `json:"Value,omitempty"`
}
