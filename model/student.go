package model

type Student struct {
	Number     int    `json:"number"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Profession string `json:"profession"`
	Reason     string `json:"reason"`
}
