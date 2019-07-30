package routes

import (
	"net/http"

	"github.com/go-macaron/session"
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

func Search(ctx *macaron.Context) {
	ctx.Params
}
