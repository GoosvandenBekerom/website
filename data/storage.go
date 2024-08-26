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

func (s *Storage) GetExperiences(ctx context.Context) ([]models.Experience, error) {
	return []models.Experience{
		{
			TimeFrom: "July 2024",
			TimeTo:   "Present",
			Company:  "Bol",
			JobTitle: "Expert Software Engineer",
			Description: "Currently working in the \"Garage\", where we aim to quickly set up and verify new business models for the company. " +
				"I really enjoy this kind of fast paced environment where there is always something different to do and you never know what the next day brings",
			Keywords: []string{"Google Cloud Platform", "Kubernetes", "Kotlin", "Postgres", "Google Cloud Pubsub", "GraphQL", "React/NextJS"},
		},
		{
			TimeFrom: "October 2023",
			TimeTo:   "June 2024",
			Company:  "Bol",
			JobTitle: "Expert Software Engineer",
			Description: "As part of the \"Product Catalog\" product, with the whole engineering team we are looking to streamline how we process content " +
				"from our suppliers to become meaningful products that you can browse on our platform. Redefining our api boundaries through domain driven design, " +
				"in order to make sure our stakeholders can continue innovating at scale.",
			Keywords: []string{"Google Cloud Platform", "Kubernetes", "Golang", "Google Cloud Bigtable", "Google Cloud Pubsub", "Postgres", "React/NextJS"},
		},
		{
			TimeFrom: "July 2023",
			TimeTo:   "September 2023",
			Company:  "Bol",
			JobTitle: "Senior Software Engineer",
			Description: "While on the look for my next long term project, I was approached by the Offer department. " +
				"They were struggling to meet deadlines for the company wide cloud migration project. " +
				"Their serving component (the offer api) is the busiest customer-facing backend service in the company, " +
				"they wanted to rebuild this service in a cloud native way using Go, and were looking for an experienced Go engineer to help out. " +
				"I agreed, under to condition that I could be the lead engineer for this project. " +
				"They agreed, so we assembled a temporary task force team. In just over 2 months, we re-designed the complete application which is now, " +
				"leveraging just Go and Bigtable, serving 100k+ rps with a p95 latency of under 20ms.",
			Keywords: []string{"Google Cloud Platform", "Kubernetes", "Golang", "Istio", "Google Cloud Bigtable", "Google Cloud Pubsub", "CAP Theorem"},
		},
		{
			TimeFrom: "October 2020",
			TimeTo:   "June 2023",
			Company:  "Bol",
			JobTitle: "Medior/Senior Software Engineer",
			Description: "I joined a new team in the Search department as a medior software engineer. " +
				"We were tasked with the difficult job of replacing an old off-the-shelf Oracle product which simply did not fit anymore considering bol.com's growth ambitions. " +
				"The company was in the middle of changing from being a retailer to being a retail platform, aiming to support tens of thousands of retailers selling their products on our shop. " +
				"We were also in the middle of migrating from our own datacenter towards the Google Cloud Platform. " +
				"In this project we fully redesigned the process of how data related to products from 25+ different data sources is prepared for the search index, " +
				"providing tooling for shop maintainers to configure this system and while keeping the original setup running with 0 downtime. " +
				"Relatively early in this project the lead engineer left the company, I took over the leading role " +
				"and somewhere along the way I got promoted to senior software engineer.",
			Keywords: []string{"Google Cloud Platform", "Kubernetes", "Golang", "Google Cloud Bigtable", "Google Cloud Pubsub", "CAP Theorem", "Postgres", "React/NextJS"},
		},
		{
			TimeFrom: "March 2019",
			TimeTo:   "Semptember 2020",
			Company:  "Bol",
			JobTitle: "Junior Software Engineer (Young Professional)",
			Description: "Fresh out of University, I joined the bol.com IT Young Professional program, a kind of traineeship that helps you kickstart your career. " +
				"I joined a new team in the Payments department, responsible for getting everything related to giftcards up to modern standards. " +
				"The company was just starting out with a migration from an on premise datacenter to the Google Cloud Platform. This was an awesome team to start in, " +
				"because we had a very clear goal and relatively small scope. This allowed us to be frontrunners in this cloud migration journey and I was able to get a lot of " +
				"experience in different technologies used throughout the company. We were the first team to have all their data available for analysis in " +
				"Google Cloud Bigquery. We also deployed the first business-critical service to production in our GCP Kubernetes setup.",
			Keywords: []string{"Google Cloud Platform", "Kubernetes", "Kotlin", "Google Cloud Pubsub", "Postgres", "Cucumber", "Angular"},
		},
	}, nil
}
