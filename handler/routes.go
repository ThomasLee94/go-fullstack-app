package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	trainers.POST("", h.CreateArticle)
	trainers.GET("/feed", h.Feed)
	trainers.PUT("/:slug", h.UpdateArticle)
	trainers.DELETE("/:slug", h.DeleteArticle)
	pokemon.POST("/:slug/comments", h.AddComment)
	pokemon.DELETE("/:slug/comments/:id", h.DeleteComment)
	pokemon.POST("/:slug/favorite", h.Favorite)
	pokemon.DELETE("/:slug/favorite", h.Unfavorite)

}
