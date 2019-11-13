package pokemon

import (
	"net/http"
	"strconv"

	"golang-starter-pack/model"
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

func (h *Handler) Articles(c echo.Context) error {
	tag := c.QueryParam("tag")
	author := c.QueryParam("author")
	favoritedBy := c.QueryParam("favorited")
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	var articles []model.Article
	var count int
	if tag != "" {
		articles, count, err = h.articleStore.ListByTag(tag, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else if author != "" {
		articles, count, err = h.articleStore.ListByAuthor(author, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else if favoritedBy != "" {
		articles, count, err = h.articleStore.ListByWhoFavorited(favoritedBy, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else {
		articles, count, err = h.articleStore.List(offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}
	return c.JSON(http.StatusOK, newArticleListResponse(h.userStore, userIDFromToken(c), articles, count))
}

func (h *Handler) Feed(c echo.Context) error {
	var articles []model.Article
	var count int
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	articles, count, err = h.articleStore.ListFeed(userIDFromToken(c), offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, newArticleListResponse(h.userStore, userIDFromToken(c), articles, count))
}
