package domain

type UserProfile struct {
	UserId  string
	Name    string
	Profile string
}

type AuthUser struct {
	UserId   string
	Email    string
	Password string
}
