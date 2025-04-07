package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/2acsek/bookings/pkg/config"
	"github.com/2acsek/bookings/pkg/handlers"
	"github.com/2acsek/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
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
	repo := handlers.NewRepo(&app)

	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	// Anonymous function
	http.HandleFunc("/anonymous", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "This endpoint is handled by an anonymous func.")

		fmt.Printf("Number of bytes written: %d\n", n)

		if err != nil {
			fmt.Println(err)
		}
	})

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal((err))
}
