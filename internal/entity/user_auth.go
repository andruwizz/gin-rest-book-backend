package entity

type UserAuth struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
