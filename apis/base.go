package apis

import (
	"github.com/go-chi/chi/v5"
	"github.com/openjogd/sodagar/core"
)

func InitApi(app core.App) (*chi.Mux, error) {
	r := chi.NewRouter()

	return r, nil
}
