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
	tags := a.Tags
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		return err
	}
	for _, t := range a.Tags {
		err := tx.Where(&model.Tag{Tag: t.Tag}).First(&t).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(&a).Association("Tags").Append(t).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Where(a.ID).Preload("Favorites").Preload("Tags").Preload("Author").Find(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	a.Tags = tags
	return tx.Commit().Error
}

func (as *PokemonStore) UpdatePokemon(a *model.Pokemon, tagList []string) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		return err
	}
	tags := make([]model.Tag, 0)
	for _, t := range tagList {
		tag := model.Tag{Tag: t}
		err := tx.Where(&tag).First(&tag).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		tags = append(tags, tag)
	}
	if err := tx.Model(a).Association("Tags").Replace(tags).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(a.ID).Preload("Favorites").Preload("Tags").Preload("Author").Find(a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *PokemonStore) DeleteArticle(a *model.Pokemon) error {
	return as.db.Delete(a).Error
}
