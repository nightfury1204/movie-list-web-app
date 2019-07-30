package models

import (
	"sync"

	"github.com/nightfury1204/movie-search-app/pkg/omdb"
)

var (
	inmem = map[string]map[string]omdb.MovieItem{}
	lock  sync.Mutex
)

func AddToMyMovieList(userID string, itm omdb.MovieItem) {
	lock.Lock()
	defer lock.Unlock()
	if _, exists := inmem[userID]; exists {
		inmem[userID][itm.ImdbID] = itm
	} else {
		inmem[userID] = map[string]omdb.MovieItem{
			itm.ImdbID: itm,
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

func GetMyMovieList(userID string) []omdb.MovieItem {
	lock.Lock()
	defer lock.Unlock()
	mlist := []omdb.MovieItem{}
	if movieList, exists := inmem[userID]; exists {
		for _, itm := range movieList {
			mlist = append(mlist, itm)
		}
	}
	return mlist
}
