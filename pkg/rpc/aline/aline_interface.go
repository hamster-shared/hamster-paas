package aline

type IUserService interface {
	GetUserByToken(token string) (User, error)
}
