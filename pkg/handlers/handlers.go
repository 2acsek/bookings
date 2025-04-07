package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/2acsek/bookings/pkg/config"
	"github.com/2acsek/bookings/pkg/models"
	"github.com/2acsek/bookings/pkg/render"
)

// The repository used by the handlers
var Repo *Repository

// The type for repositories
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	stringMap := map[string]string{
		"color": "#ABC123",
	}
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	n, err := fmt.Fprintf(w, "Your IP is: %s", remoteIP)

	fmt.Printf("Number of bytes written: %d\n", n)

	if err != nil {
		fmt.Println(err)
	}
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	x := 10.0
	y := 0.0
	f, err := divideValues(10, 0)
	if err != nil {
		fmt.Fprintf(w, "%s", fmt.Sprintf("Error: %s\n", err.Error()))
	} else {
		fmt.Fprintf(w, "%s", fmt.Sprintf("%f is divided by %f is %f", x, y, f))
	}
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}
	return x / y, nil
}
