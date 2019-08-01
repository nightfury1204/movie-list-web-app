package models

import (
	"sync"

	"github.com/nightfury1204/movie-listing-app/pkg/omdb"
)

var (
	inmem = map[string]map[string]omdb.MovieDetails{}
	lock  sync.Mutex
)

func AddToMyMovieList(userID string, movieDetails omdb.MovieDetails) {
	lock.Lock()
	defer lock.Unlock()
	if _, exists := inmem[userID]; exists {
		inmem[userID][movieDetails.ImdbID] = movieDetails
	} else {
		inmem[userID] = map[string]omdb.MovieDetails{
			movieDetails.ImdbID: movieDetails,
		}
	}
}

func RemoveFromMyMovieList(userID string, imdbID string) {
	lock.Lock()
	defer lock.Unlock()
	if movieList, exists := inmem[userID]; exists {
		if _, has := movieList[imdbID]; has {
			delete(movieList, imdbID)
		}
	}
}

func GetMyMovieList(userID string) []omdb.MovieDetails {
	lock.Lock()
	defer lock.Unlock()
	mlist := []omdb.MovieDetails{}
	if movieList, exists := inmem[userID]; exists {
		for _, itm := range movieList {
			mlist = append(mlist, itm)
		}
	}
	return mlist
}
