package model

type UserCreateInput struct {
	Password string
	Passpost string
	Nickname string
}

type UserSignInInput struct {
	Passport string
	Password string
}
