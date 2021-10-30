package repository

import (
	"errors"
	"survey/api/middlewares"
	"survey/api/models"
	"survey/api/models/dto"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	SaveUser(user *models.User) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
	FindUserByID(uid string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(uid string) (string, error)
	LoginByUsername(dtoLogin *dto.Login) (string, error)
	UnActiveUser(id string) (*models.User, error)
	ActivatedUser(id string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) SaveUser(newUser *models.User) (*models.User, error) {
	var err error
	err = u.db.Debug().Create(&newUser).Error
	if err != nil {
		return &models.User{}, err
	}
	return newUser, nil
}

func (u *userRepository) FindAllUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	err = u.db.Debug().Model(&models.User{}).Limit(100).Error
	if err != nil {
		return &users, err
	}
	return &users, nil
}

func (u *userRepository) FindUserByID(id string) (*models.User, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var users models.User
	err = u.db.Debug().Table("user").Where("id  = ?", uid).Error
	if err != nil {
		return &users, err
	}
	if gorm.ErrRecordNotFound == err {
		return &models.User{}, errors.New("User Not Found")
	}
	return &users, err
}

func (u *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	u.db = u.db.Debug().Save(user)
	if err := u.db.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) DeleteUser(uid string) (string, error) {
	u.db = u.db.Debug().Model(&models.User{}).Where("id = ?", uid).Delete(&models.User{})

	if u.db.Error != nil {
		return "", u.db.Error

	}
	return string(rune(u.db.RowsAffected)), nil
}

func (a *userRepository) LoginByUsername(dtoLogin *dto.Login) (string, error) {

	var err error
	var users models.User

	if err = a.db.Debug().Table("users").Where("username = ?", dtoLogin.Username).Preload(clause.Associations).Find(&users).Error; err != nil {
		return "", err
	}
	err = models.VerifyPassword(users.Password, dtoLogin.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return middlewares.CreateToken(&users)
}

func (u *userRepository) UnActiveUser(id string) (*models.User, error) {
	uid, _ := uuid.FromString(id)
	var users models.User
	u.db = u.db.Debug().Model(&users).Where("id = ?", uid).UpdateColumns(
		map[string]interface{}{
			"is_active":  false,
			"deleted_at": time.Now(),
		},
	)
	if err := u.db.Error; err != nil {
		return &users, err
	}
	return &users, nil
}

func (u *userRepository) ActivatedUser(id string) (*models.User, error) {
	uid, _ := uuid.FromString(id)
	var users models.User
	u.db = u.db.Debug().Model(&users).Where("id = ?", uid).UpdateColumns(
		map[string]interface{}{
			"is_active":  true,
			"deleted_at": time.Now(),
		},
	)
	if err := u.db.Error; err != nil {
		return &users, err
	}
	return &users, nil
}
