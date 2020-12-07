package main

/**
    ===================
	Author: Mark Volkov
    ===================
 */

import (
	"context"
	"efieldrestful/app"
	"flag"
	"log"
	"time"
)

const (
	TIMEOUT = 10
)

func main() {
	envFlag := flag.String("env", "dev", "Your environment config to run: ( dev || prod )")
	flag.Parse()
	app := app.App{}
	app.Init(*envFlag)
	log.Println(*envFlag)
	app.RunApplication()
	defer func() {
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
		if err := app.DatabaseService.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
