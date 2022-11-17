package route

import (
	"fmt"
	"net/http"
	activitygroups "todo-list/internal/modules/activity-groups"

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

}
