package testdata

type UserRepository struct {
}

func (r *UserRepository) GetByID(id string) (*User, error) {
	user := &User{Name: "John Doe"}
	return user, nil
}
