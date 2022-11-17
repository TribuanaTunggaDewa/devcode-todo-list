package activitygroups

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetById)
	g.POST("", h.Store)
	g.PATCH("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
