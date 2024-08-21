package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/goosvandenbekerom/website/data/models"
	// Sqlite db driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrNotFound = errors.New("not found")
)

const dbFileEnvVar = "DB_FILENAME"

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	file, set := os.LookupEnv(dbFileEnvVar)
	if !set {
		return nil, fmt.Errorf("%s env variable not set", dbFileEnvVar)
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if err := applyMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %s", err)
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) GetProfile(ctx context.Context) (models.Profile, error) {
	row := s.db.QueryRowContext(ctx, "SELECT data FROM json_data WHERE key='profile'")
	if err := row.Err(); err != nil {
		slog.Error("db error", slog.String("error", err.Error()))
		return models.Profile{}, fmt.Errorf("profile: %w", ErrNotFound)
	}

	var data string
	if err := row.Scan(&data); err != nil {
		return models.Profile{}, fmt.Errorf("failed to read profile data from row: %s", err)
	}

	var profile models.Profile
	if err := json.Unmarshal(json.RawMessage(data), &profile); err != nil {
		return models.Profile{}, fmt.Errorf("failed to parse profile data: %s", err)
	}

	return profile, nil
}
