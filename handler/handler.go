package handler

import (
	"golang-starter-pack/pokemon"
	"golang-starter-pack/trainer"
)

type Handler struct {
	pokemonStore pokemon.Store
	trainerStore trainer.Store
}

func NewHandler(us pokemon.Store, as trainer.Store) *Handler {
	return &Handler{
		pokemonStore: us,
		trainerStore: as,
	}
}
