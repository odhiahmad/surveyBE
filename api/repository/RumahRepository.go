package repository

import (
	"errors"
	"survey/api/models"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IRumahRepository interface {
	SaveRumah(rumah *models.Rumah) (*models.Rumah, error)
	FindRumahByID(uid string) (*models.Rumah, error)
	UpdateARumah(rumah *models.Rumah) (*models.Rumah, error)
	DeleteARumah(uid string) (string, error)
	UnActiveRumah(id string) (*models.Rumah, error)
	ActivatedRumah(id string) (*models.Rumah, error)
	GetAllRumah(page, pageSize, order string) ([]*models.Rumah, error)
}

type rumahRepository struct {
	db *gorm.DB
}

func NewRumahRepository(db *gorm.DB) IRumahRepository {
	return &rumahRepository{
		db,
	}
}

func (r *rumahRepository) SaveRumah(newRumah *models.Rumah) (*models.Rumah, error) {
	var err error
	err = r.db.Debug().Create(&newRumah).Error
	if err != nil {
		return &models.Rumah{}, err
	}
	return newRumah, nil
}

func (r *rumahRepository) FindRumahByID(id string) (*models.Rumah, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var rumahs models.Rumah
	err = r.db.Debug().Table("rumah").Where("id  = ?", uid).Error
	if err != nil {
		return &rumahs, err
	}
	if gorm.ErrRecordNotFound == err {
		return &models.Rumah{}, errors.New("Rumah Not Found")
	}
	return &rumahs, err
}

func (r *rumahRepository) UpdateARumah(rumah *models.Rumah) (*models.Rumah, error) {
	r.db = r.db.Debug().Save(rumah)
	if err := r.db.Error; err != nil {
		return rumah, err
	}
	return rumah, nil
}

func (r *rumahRepository) DeleteARumah(uid string) (string, error) {
	r.db = r.db.Debug().Model(&models.Rumah{}).Where("id = ?", uid).Delete(&models.Rumah{})

	if r.db.Error != nil {
		return "", r.db.Error

	}
	return string(rune(r.db.RowsAffected)), nil
}

func (r *rumahRepository) UnActiveRumah(id string) (*models.Rumah, error) {
	uid, _ := uuid.FromString(id)
	var rumahs models.Rumah
	r.db = r.db.Debug().Model(&rumahs).Where("id = ?", uid).UpdateColumns(
		map[string]interface{}{
			"is_active":  false,
			"deleted_at": time.Now(),
		},
	)
	if err := r.db.Error; err != nil {
		return &rumahs, err
	}
	return &rumahs, nil
}

func (r *rumahRepository) ActivatedRumah(id string) (*models.Rumah, error) {
	uid, _ := uuid.FromString(id)
	var rumahs models.Rumah
	r.db = r.db.Debug().Model(&rumahs).Where("id = ?", uid).UpdateColumns(
		map[string]interface{}{
			"is_active":  true,
			"deleted_at": time.Now(),
		},
	)
	if err := r.db.Error; err != nil {
		return &rumahs, err
	}
	return &rumahs, nil
}

func (w *rumahRepository) GetAllRumah(page, pageSize, order string) ([]*models.Rumah, error) {
	rumahs := make([]*models.Rumah, 0)
	if order != "" {
		if err := w.db.Debug().Scopes(Paginate(page, pageSize)).First(&rumahs).Limit(3).Error; err != nil {
			return nil, err
		}
		return rumahs, nil
	} else {
		if err := w.db.Debug().Scopes(Paginate(page, pageSize)).Order("created_at desc").Find(&rumahs).Limit(3).Error; err != nil {
			return nil, err
		}
		return rumahs, nil
	}

}
