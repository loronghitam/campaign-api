package usecase

import "campaign-api/domain"

type UserFormat struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Occupation string `json:"occupation,omitempty"`
	Email      string `json:"email,omitempty"`
	Token      string `json:"token,omitempty"`
}

func NewFormatUser(user domain.User, token string) UserFormat {
	format := UserFormat{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return format
}
