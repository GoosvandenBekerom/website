package data

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/goosvandenbekerom/website/data/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type Storage struct {
	db map[string]json.RawMessage
}

func NewStorage() *Storage {
	return &Storage{
		db: map[string]json.RawMessage{
			"profile": json.RawMessage(`{"name": "Goos van den Bekerom", "email": "goos.bekerom@gmail.com", "date_of_birth": "1995-06-07T00:00:00Z"}`),
		},
	}
}

func (s *Storage) GetProfile() (models.Profile, error) {
	raw, found := s.db["profile"]
	if !found {
		return models.Profile{}, fmt.Errorf("profile: %w", ErrNotFound)
	}

	var profile models.Profile
	if err := json.Unmarshal(raw, &profile); err != nil {
		return models.Profile{}, fmt.Errorf("failed to parse profile data: %s", err)
	}

	return profile, nil
}
