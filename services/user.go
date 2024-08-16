package greetings

type UserService struct {
	db string
}

func (s UserService) GetUserName() string {
	return "Bob"
}
