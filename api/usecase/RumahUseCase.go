package usecase

import (
	"survey/api/models"
	"survey/api/repository"
)

type IRumahUseCase interface {
	SaveRumah(rumah *models.Rumah) (*models.Rumah, error)
	UpdateInfo(editRumah *models.Rumah) (*models.Rumah, error)
	Unregister(id string) (string, error)
	FindRumahById(id string) (*models.Rumah, error)
	GetAllRumah(page, pageSize, order string) ([]*models.Rumah, error)
}

type RumahUseCaseRepo struct {
	rumahRepo repository.IRumahRepository
}

func (u *RumahUseCaseRepo) SaveRumah(rumah *models.Rumah) (*models.Rumah, error) {
	return u.rumahRepo.SaveRumah(rumah)
}
func (u *RumahUseCaseRepo) FindRumahById(id string) (*models.Rumah, error) {
	return u.rumahRepo.FindRumahByID(id)
}

func (u *RumahUseCaseRepo) Unregister(id string) (string, error) {
	return u.rumahRepo.DeleteARumah(id)
}

func (u *RumahUseCaseRepo) UpdateInfo(editRumah *models.Rumah) (*models.Rumah, error) {
	return u.rumahRepo.UpdateARumah(editRumah)
}

func (u *RumahUseCaseRepo) GetAllRumah(page, pageSize, order string) ([]*models.Rumah, error) {
	return u.rumahRepo.GetAllRumah(page, pageSize, order)
}

func NewRumahUseCase(rumahRepo repository.IRumahRepository) IRumahUseCase {
	return &RumahUseCaseRepo{
		rumahRepo,
	}
}
