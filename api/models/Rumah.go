package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Rumah struct {
	ID                 uuid.UUID  `gorm:"type:uuid;unique;index" json:"id"`
	Kecamatan          string     `gorm:"unique;not null; size: 255" json:"kecamatan"`
	Desa               string     `gorm:"not null; size: 255" json:"desa"`
	Nagari             string     `gorm:"not null; size: 255" json:"nagari"`
	Jorong             string     `gorm:"not null; size: 255" json:"jorong"`
	Dusun              string     `gorm:"not null; size: 255" json:"dusun"`
	Rt                 string     `gorm:"not null; size: 255" json:"rt"`
	NomorRumah         string     `gorm:"not null; size: 255" json:"nomor_rumah"`
	Lat                string     `gorm:"not null; size: 255" json:"lat"`
	Long               string     `gorm:"not null; size: 255" json:"long"`
	JumlahKeluarga     int32      `gorm:"not null" json:"jumlah_keluarga"`
	JumlahPenghuni     int32      `gorm:"not null" json:"jumlah_penghuni"`
	NamaKepalaKeluarga string     `gorm:"not null; size: 255" json:"nama_kepala_keluarga"`
	NomorKK            string     `gorm:"not null; size: 255" json:"nomor_kk"`
	StatusKepemilikan  int8       `gorm:"not null" json:"status_kepemilikan"`
	LuasRumah          int64      `gorm:"not null" json:"luas_rumah"`
	Kondisi            int8       `gorm:"not null" json:"kondisi"`
	Jenis              int8       `gorm:"not null" json:"jenis"`
	IsActive           bool       `gorm:"not null; column:is_active"`
	CreatedAt          time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (u *Rumah) Prepare() error {
	u.ID = uuid.NewV4()

	return nil
}
