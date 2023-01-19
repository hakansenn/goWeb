package handlers

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/hakansenn/goWeb/internals/config"
	"github.com/hakansenn/goWeb/internals/models"
	"github.com/hakansenn/goWeb/internals/render"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes () http.Handler{
//what am i going to put in the session
gob.Register(models.Reservation{})

app.InProduction = false

session = scs.New()
session.Lifetime = 24 * time.Hour
session.Cookie.Persist = true
session.Cookie.SameSite = http.SameSiteLaxMode
session.Cookie.Secure = app.InProduction

app.Session = session

tc, err := render.CreateTemplateCache()
if err != nil {
	log.Fatal("cannot create template cache")
}

app.TemplateCache = tc
app.UseCache = false

repo := NewRepo(&app)
NewHandlers(repo)
render.NewTemplates(&app)

mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
