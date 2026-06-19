package testdata

type UserController struct {
	service UserService
}

func (c *UserController) GetUser(id string) (*User, error) {
	user, err := c.service.GetUser(id)
	return user, err
}
