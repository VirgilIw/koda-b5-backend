package dto

type Authentication struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Register struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
