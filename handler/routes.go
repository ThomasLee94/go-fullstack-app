package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	trainers.POST("/:trainer", h.CreateTrainer)
	trainers.GET("/:trainer", h.GetTrainer)
	trainers.PUT("/:trainer", h.UpdateTrainer)
	trainers.DELETE("/:trainer", h.DeleteTrainer)
	pokemon.POST("/:pokemon", h.AddPokemon)
	pokemon.DELETE("/:pokemon", h.DeletePokemon)
	pokemon.POST("/:pokemon", h.GetPokemon)

}
