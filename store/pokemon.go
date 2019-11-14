package store

import (
	"golang-starter-pack/model"

	"github.com/jinzhu/gorm"
)

type PokemonStore struct {
	db *gorm.DB
}

func NewPokemonStore(db *gorm.DB) *PokemonStore {
	return &PokemonStore{
		db: db,
	}
}

func (as *PokemonStore) CreatePokemon(a *model.Pokemon) error {
	// create Pokemon
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		return err
	}
}

func (as *PokemonStore) UpdatePokemon(a *model.Pokemon, tagList []string) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		return err
	}
}

func (as *PokemonStore) DeleteArticle(a *model.Pokemon) error {
	return as.db.Delete(a).Error
}
