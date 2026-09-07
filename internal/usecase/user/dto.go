package user

type UserCreateParam struct {
	Name     string
	Email    string
	Password string
}

type UserLoginParam struct {
	Email    string
	Password string
}

type UserGetParam struct {
	Name  string
	Email string
}
