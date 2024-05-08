package models

type TeamData struct {
	ID       int    `json:"id"`
	TeamName string `json:"team_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DataGame string `json:"data_game"`
	NameGame string `json:"name_game"`
}
