package testdata

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUser(id string) (*User, error) {
	user, err := s.repo.GetByID(id)
	return user, err
}
