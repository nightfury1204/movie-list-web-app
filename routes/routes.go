package routes

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-macaron/session"
	"github.com/nightfury1204/movie-search-app/pkg/logger"
	"github.com/nightfury1204/movie-search-app/pkg/omdb"
	"gopkg.in/macaron.v1"
)

const (
	SessionIDKey = "session_id"
	UserIDKey    = "user_id"
)

// NewMacaron initializes Macaron instance.
func NewMacaron() *macaron.Macaron {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	return m
}

func RegisterRoutes(m *macaron.Macaron) {
	m.Get("/", Authenticate, Home)
	m.Get("/search", Authenticate, Search)
	m.Get("/movie", Authenticate, Movie)
}

func Authenticate(sess session.Store, ctx *macaron.Context) {
	if sessionID, ok := sess.Get(SessionIDKey).(string); !ok || sessionID == "" {
		ctx.Redirect("/login")
	} else {
		// currently user id is set as current id
		ctx.Data[UserIDKey] = sessionID
	}
}

func Home(ctx *macaron.Context) {
	ctx.HTML(http.StatusOK, "home")
}

// Search will fetch the movie search results
func Search(ctx *macaron.Context) {
	log := logger.GetLogger()
	log = log.WithValues(UserIDKey, ctx.Data[UserIDKey].(string), "operation", "search")

	s := ctx.Req.URL.Query().Get("s")
	if len(s) == 0 {
		err := errors.New("search value is not provided")
		log.Error(err)
		ctx.HTML(http.StatusBadRequest, "search_error", map[string]string{
			"error":   err.Error(),
			"keyword": s,
		})
	}

	page := ctx.Req.URL.Query().Get("page")
	if len(pageNo) == 0 {
		page = "1"
	}

	searchResp, status, err := omdb.SearchMovie(s, pageNo)
	if searchResp != nil && searchResp.Error != "" {
		status = http.StatusBadRequest
		err = errors.New(searchResp.Error)
	}
	if err != nil {
		log.Error(err)
		ctx.HTML(status, "search_error", map[string]string{
			"error":   err.Error(),
			"keyword": s,
		})
	}

	ctx.HTML(http.StatusOK, "search_result", searchResp)
}

// Movie will fetch the movie details
func Movie(ctx *macaron.Context) {
	log := logger.GetLogger()
	log = log.WithValues(UserIDKey, ctx.Data[UserIDKey].(string), "operation", "movie details")

	imdbID := ctx.Req.URL.Query().Get("i")
	if len(imdbID) == 0 {
		err := errors.New("imdb id is not provided")
		log.Error(err)
		ctx.HTML(http.StatusBadRequest, "movie_details_error", map[string]string{
			"error": err.Error(),
		})
	}

	movieDetails, status, err := omdb.GetMovieDetails(imdbID)
	if movieDetails != nil && movieDetails.Error != "" {
		status = http.StatusBadRequest
		err = errors.New(movieDetails.Error)
	}
	if err != nil {
		log.Error(err))
		ctx.HTML(http.StatusBadRequest, "movie_details_error", map[string]string{
			"error": err.Error(),
		})
	}

	ctx.HTML(http.StatusOK, "movie_details", movieDetails)
}
