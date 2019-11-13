package pokemon

import (
	"golang-starter-pack/model"
)

type Store interface {
	CreatePokemon(*model.Pokemon) error
	UpdatePokemon(*model.Pokemon, []string) error
	DeletePokemon(*model.Pokemon) error

	CreateTrainer(*model.Trainer) error
	UpdateTrainer(*model.Pokemon, []string) error
	DeleteTrainer(*model.Pokemon) error
}
