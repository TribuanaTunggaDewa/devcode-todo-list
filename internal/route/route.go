package route

import (
	"fmt"
	"net/http"
	activitygroups "todo-list/internal/modules/activity-groups"
	todoitems "todo-list/internal/modules/todo-items"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {

	var (
		APP     = "Todo-App"
		VERSION = "0.0.1"
	)

	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	//routes
	activitygroups.NewHandler().Route(g.Group("activity-groups"))
	todoitems.NewHandler().Route(g.Group("todo-items"))
}
