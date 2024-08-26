package models

type Experience struct {
	TimeFrom    string   `json:"time_from"`
	TimeTo      string   `json:"time_to"`
	Company     string   `json:"company"`
	JobTitle    string   `json:"job_title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}
