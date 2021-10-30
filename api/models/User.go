package models

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;unique;index" json:"id"`
	Username  string     `gorm:"unique;not null; size: 255" json:"username"`
	Password  string     `gorm:"not null; size: 50" json:"password"`
	IsActive  bool       `gorm:"not null; column:is_active"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

type UserRequest struct {
	Username string
	Password string
	Name     string
	IsActive bool
}

type UserResponse struct {
	ID        uuid.UUID
	Name      string
	Username  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func (u *User) Prepare() error {
	u.ID = uuid.NewV4()
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.IsActive = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":

		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Minimum eight characters, at least one letter and one number")
		}

		return nil
	case "forgot":
		if u.Password == "" {
			return errors.New("Minimum eight characters, at least one letter and one number")
		}
		return nil
	default:
		if u.Username == "" {
			return errors.New("Required username")
		}
		if u.Password == "" {
			return errors.New("Minimum eight characters, at least one letter and one number")
		}
		return nil
	}
}
