package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	var user User
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repo.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	_, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	return true, nil
}
