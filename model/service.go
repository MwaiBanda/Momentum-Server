package model

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type		string `json:"type"`
	Image       string `json:"image"`
}	