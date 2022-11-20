package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-list/database"
	"todo-list/internal/route"
	"todo-list/pkg/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	env := util.NewEnv()
	env.Load()
}

func main() {

	exit := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	fmt.Print("init app")

	database.Init(ctx)

	e := echo.New()

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | &{latency_human} | ${remote_ip} | ${method} | ${uri}", "Todo-App"),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)

	e.Validator = &util.CustomValidator{Validator: validator.New()}

	route.Init(e.Group(""))

	go func() {
		err := e.Start(":3030")
		if err != http.ErrServerClosed {
			e.Logger.Fatal()
		}
	}()

	<-exit
	cancel()
	ctx, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	e.Shutdown(ctx)

}
