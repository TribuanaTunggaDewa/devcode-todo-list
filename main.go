package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"todo-list/database"
	"todo-list/pkg/util"
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

}
