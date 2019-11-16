package handler

import (
	"golang-starter-pack/model"
)

type pokemonResponse struct {
	Pokemon struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"pokemon"`
}

func newPokemonResponse(u *model.Pokemon) *pokemonResponse {
	r := new(pokemonResponse)
	r.Pokemon.Name = u.Name
	r.Pokemon.Type = u.Type
	return r
}

type trainerResponse struct {
	Trainer struct {
		Name    string `json:"name"`
		Pokemon []Pokemon
	} `json:"pokemon"`
}

func newTrainerResponse(u *model.Trainer) *trainerResponse {
	r := new(trainerResponse)
	r.Trainer.Name = u.Name
	r.Trainer.Pokemon = u.Pokemon
	return r
}
