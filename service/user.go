package service

import "go-nornia/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) VerifyUser(tel string) (bool, error) {
	user := models.NewUserModel()
	if err := user.GetByTel(tel); err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}
