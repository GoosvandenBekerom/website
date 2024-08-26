package handlers

import (
	"net/http"

	"github.com/goosvandenbekerom/website/data"
	"github.com/goosvandenbekerom/website/web/internal/templates/pages"
)

type experience struct {
	storage *data.Storage
}

func Experience(storage *data.Storage) experience {
	return experience{
		storage: storage,
	}
}

func (h experience) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.storage.GetExperiences(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := pages.Experience(experiences).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
