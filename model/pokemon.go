package model

import "github.com/jinzhu/gorm"

type Pokemon struct {
	gorm.Model
	Slug      string `gorm:"unique_index;not null"`
	Name      string `gorm:"not null"`
	Type      string
	PokemonId uint
}
