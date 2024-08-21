package handlers

import (
	"net/http"

	"github.com/goosvandenbekerom/website/data"
	"github.com/goosvandenbekerom/website/web/internal/templates/pages"
)

type home struct {
	storage *data.Storage
}

func Home(storage *data.Storage) home {
	return home{
		storage: storage,
	}
}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	profile, err := h.storage.GetProfile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := pages.Home(profile).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
