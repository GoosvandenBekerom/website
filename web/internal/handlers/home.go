package handlers

import (
	"net/http"

	"github.com/goosvandenbekerom/website/web/internal/templates/pages"
)

type home struct {
}

func Home() home {
	return home{}
}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := pages.Home().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
