package service

type User struct {
	Name string
}

func GetUser(id string) (*User, error) {

	user, err := repo.GetByID(id)
	return user, err
}
