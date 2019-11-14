package pokemon

import (
	"net/http"

	"golang-starter-pack/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPokemon(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.articleStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newPokemonResponse(c, a))
}
