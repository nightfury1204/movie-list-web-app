package routes

import (
	"encoding/json"
	"github.com/nightfury1204/movie-search-app/models"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"github.com/nightfury1204/movie-search-app/pkg/logger"
	"github.com/nightfury1204/movie-search-app/pkg/omdb"
	"gopkg.in/macaron.v1"
)

const (
	SessionIDKey = "session_id"
	UserIDKey    = "user_id"
)

type LoginForm struct {
	UserID   string `form:"user_id" binding:"Required"`
	Password string `form:"password" binding:"Required"`
}

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
	m.Get("/movie/details", Authenticate, MovieDetails)
	
	m.Group("/mylist/movies", func() {
		m.Get("", GetMyMovieList)
		m.Post("/:id", AddToMyMovieList)
		m.Delete("/:id", RemoveFromMyMovieList)
	}, Authenticate)

	m.Group("/login", func() {
		m.Post("", binding.Bind(LoginForm{}), Login)
		m.Get("", func(ctx *macaron.Context) {
			ctx.HTML(http.StatusOK, "login")
		})
	})
	m.Post("/login", binding.Bind(LoginForm{}), Login)
	m.Get("/logout", Logout)
}

func Logout(sess session.Store, ctx *macaron.Context) {
	sess.Destory(ctx)
	ctx.Redirect("/login")
}

func Login(login LoginForm, sess session.Store, ctx *macaron.Context) {
	var userID string
	if login.UserID == "user1" && login.Password == "pass1" {
		userID = "user1"
	} else if login.UserID == "user2" && login.Password == "pass2" {
		userID = "user2"
	} else if login.UserID == "user3" && login.Password == "pass3" {
		userID = "user3"
	} else {
		ctx.HTML(http.StatusUnauthorized, "login", map[string]string{
			"error": "invalid userid/password",
		})
		return
	}
	sess.Set(SessionIDKey, userID)
	ctx.Redirect("/")
}

func Authenticate(sess session.Store, ctx *macaron.Context) {
	if sessionID, ok := sess.Get(SessionIDKey).(string); !ok || sessionID == "" {
		ctx.Redirect("/login")
		return
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
	log := logger.GetLogger().WithValues(UserIDKey, ctx.Data[UserIDKey].(string), "operation", "search")

	s := ctx.Req.URL.Query().Get("s")
	if len(s) == 0 {
		err := errors.New("search value is not provided")
		log.Error(err, "failed to search movie")
		ctx.HTML(http.StatusBadRequest, "search_error", map[string]string{
			"error":   err.Error(),
			"keyword": s,
		})
		return
	}

	page := ctx.Req.URL.Query().Get("page")
	if len(page) == 0 {
		page = "1"
	}

	searchResp, status, err := omdb.SearchMovie(s, page)
	if searchResp != nil && searchResp.Error != "" {
		status = http.StatusBadRequest
		err = errors.New(searchResp.Error)
	}
	if err != nil {
		log.Error(err, "failed to search movie")
		ctx.HTML(status, "search_error", map[string]string{
			"error":   err.Error(),
			"keyword": s,
		})
		return
	}

	ctx.HTML(http.StatusOK, "search_result", searchResp)
}

// MovieDetails will fetch the movie details
func MovieDetails(ctx *macaron.Context) {
	log := logger.GetLogger().WithValues(UserIDKey, ctx.Data[UserIDKey].(string), "operation", "movie details")

	imdbID := ctx.Req.URL.Query().Get("i")
	if len(imdbID) == 0 {
		err := errors.New("imdb id is not provided")
		log.Error(err, "failed to get movie details")
		ctx.HTML(http.StatusBadRequest, "movie_details_error", map[string]string{
			"error": err.Error(),
		})
		return
	}

	movieDetails, status, err := omdb.GetMovieDetails(imdbID)
	if movieDetails != nil && movieDetails.Error != "" {
		status = http.StatusBadRequest
		err = errors.New(movieDetails.Error)
	}
	if err != nil {
		log.Error(err, "failed to get movie details")
		ctx.HTML(status, "movie_details_error", map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "movie_details", movieDetails)
}

func GetMyMovieList(ctx *macaron.Context) {
	userID := ctx.Data[UserIDKey].(string)
	log := logger.GetLogger().WithValues(UserIDKey, userID, "operation", "get my movie list")

	log.V(3).Info("fetching movie list")
	movieList := models.GetMyMovieList(userID)
	ctx.HTML(http.StatusOK, "my_list", movieList)
}

func RemoveFromMyMovieList(ctx *macaron.Context)  {
	userID := ctx.Data[UserIDKey].(string)
	log := logger.GetLogger().WithValues(UserIDKey, userID, "operation", "remove from my movie list")

	imdbID := ctx.Params(":id")

	log.V(3).Info("removing movie from the list", "imdbID", imdbID)
	models.RemoveFromMyMovieList(userID, imdbID)
	ctx.Redirect("/mylist/movies", http.StatusOK)
}

func AddToMyMovieList(ctx *macaron.Context)  {
	userID := ctx.Data[UserIDKey].(string)
	log := logger.GetLogger().WithValues(UserIDKey, userID, "operation", "add to my movie list")

	item := &omdb.MovieItem{}
	r := ctx.Req.Request
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		log.Error(err, "failed to unmarshal movie item")
		ctx.JSON(http.StatusBadRequest, "add_movie_failed")
	}

	log.V(3).Info("Adding movie in the list", "imdbID", item.ImdbID)
	models.AddToMyMovieList(userID, *item)
}