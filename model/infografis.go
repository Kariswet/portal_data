package model

import (
	"encoding/json"

	"github.com/lib/pq"
)

type Infografis struct {
	ID          string          `json:"id" gorm:"primaryKey"`
	Kategori    string          `json:"kategori"`
	Judul       string          `json:"judul"`
	Publisher   string          `json:"publisher"`
	Tags        pq.StringArray  `json:"tags" gorm:"type:text[]"`
	ImageSlider json.RawMessage `json:"image_slider" gorm:"type:jsonb"`
	Deskripsi   string          `json:"deskripsi"`
	CreatedAt   int64           `json:"created_at"`
	UpdatedAt   int64           `json:"updated_at"`
}

type InfografisView struct {
	Kategori    string   `json:"kategori" example:""`
	Judul       string   `json:"judul" example:""`
	Publisher   string   `json:"publisher" example:""`
	Tags        []string `json:"tags" example:"" gorm:"type:text[]"`
	ImageSlider []struct {
		File      string `json:"file" example:""`
		Deskripsi string `json:"deskripsi" example:""`
	} `json:"image_slider"`
	Deskripsi string `json:"deskripsi" example:""`
}

// @Schema
type pqStringArray []string
