package usecase

import (
	"survey/api/models"
	"survey/api/models/dto"
	"survey/api/repository"
)

type IUserUseCase interface {
	LoginByUsername(dtoLogin *dto.Login) (string, error)
	SaveUser(user *models.User) (*models.User, error)
	UpdateInfo(editUser *models.User) (*models.User, error)
	Unregister(id string) (string, error)
	FindUserById(id string) (*models.User, error)
}

type UserUseCaseRepo struct {
	userRepo repository.IUserRepository
}

func (u *UserUseCaseRepo) LoginByUsername(dtoLogin *dto.Login) (string, error) {
	return u.userRepo.LoginByUsername(dtoLogin)
}

func (u *UserUseCaseRepo) SaveUser(user *models.User) (*models.User, error) {
	return u.userRepo.SaveUser(user)

}
func (u *UserUseCaseRepo) FindUserById(id string) (*models.User, error) {
	return u.userRepo.FindUserByID(id)
}

func (u *UserUseCaseRepo) Unregister(id string) (string, error) {
	return u.userRepo.DeleteUser(id)
}

func (u *UserUseCaseRepo) UpdateInfo(editUser *models.User) (*models.User, error) {
	return u.userRepo.UpdateUser(editUser)
}

func NewUserUseCase(userRepo repository.IUserRepository) IUserUseCase {
	return &UserUseCaseRepo{
		userRepo,
	}
}
