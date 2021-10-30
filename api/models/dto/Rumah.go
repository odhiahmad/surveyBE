package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type RumahRequest struct {
	Kecamatan          string
	Desa               string
	Nagari             string
	Jorong             string
	Dusun              string
	Rt                 string
	NomorRumah         string
	Lat                string
	Long               string
	JumlahKeluarga     int32
	JumlahPenghuni     int32
	NamaKepalaKeluarga string
	NomorKK            string
	StatusKepemilikan  int8
	LuasRumah          int64
	Kondisi            int8
	Jenis              int8
	IsActive           bool
}

type RumahResponse struct {
	ID                 uuid.UUID
	Kecamatan          string
	Desa               string
	Nagari             string
	Jorong             string
	Dusun              string
	Rt                 string
	NomorRumah         string
	Lat                string
	Long               string
	JumlahKeluarga     int32
	JumlahPenghuni     int32
	NamaKepalaKeluarga string
	NomorKK            string
	StatusKepemilikan  int8
	LuasRumah          int64
	Kondisi            int8
	Jenis              int8
	IsActive           bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
