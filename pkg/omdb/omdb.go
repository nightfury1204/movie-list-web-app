package omdb

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

var client *Client

type Client struct {
	apiKey string
	apiUrl string
}

func Initialize(url, apiKey string) {
	client = &Client{
		apiKey: apiKey,
		apiUrl: url,
	}
}

// 's' is the provided search keyword
// 'page' is the page number of the search results
func SearchMovie(s, page string) (*SearchResponse, int, error) {
	if client == nil {
		return nil, http.StatusInternalServerError, errors.New("client is not initialized")
	}

	u, err := url.Parse(client.apiUrl)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to parse omdb api url")
	}

	qParams := u.Query()
	qParams.Set("apikey", client.apiKey)
	qParams.Set("s", s)
	qParams.Set("type", "movie")
	if len(page) > 0 {
		qParams.Set("page", page)
	}
	u.RawQuery = qParams.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	searchResp := &SearchResponse{}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(searchResp); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to unmarshal search response")
	}
	return searchResp, http.StatusOK, nil
}

func GetMovieDetails(imdbID string) (*MovieDetails, int, error) {
	if client == nil {
		return nil, http.StatusInternalServerError, errors.New("client is not initialized")
	}

	u, err := url.Parse(client.apiUrl)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to parse omdb api url")
	}

	qParams := u.Query()
	qParams.Set("apikey", client.apiKey)
	qParams.Set("i", imdbID)
	qParams.Set("type", "movie")
	u.RawQuery = qParams.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	movieDetails := &MovieDetails{}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(movieDetails); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to unmarshal movie details response")
	}
	return movieDetails, http.StatusOK, nil
}
