package store

import (
	"golang-starter-pack/model"

	"github.com/jinzhu/gorm"
)

type TraierStore struct {
	db *gorm.DB
}

func NewTrainerStore(db *gorm.DB) *TrainerStore {
	return &PokemonStore{
		db: db,
	}
}

func (as *TrainerStore) CreatePokemon(a *model.Trainer) error {
	// create Pokemon
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		return err
	}
}

func (as *TrainerStore) UpdateTrainer(a *model.Trainer, tagList []string) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		return err
	}
}

func (as *TrainerStore) DeleteTrainer(a *model.Trainer) error {
	return as.db.Delete(a).Error
}
