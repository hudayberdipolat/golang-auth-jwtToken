package main

import (
	"fmt"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/app"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/setup/constructor"
	"log"
)

func main() {
	appDependencies, err := app.GetAppDependencies()
	if err != nil {
		log.Fatal("app dependencies error :--->", err.Error())
	}
	constructor.Build(appDependencies)
	appRoutes := app.NewApp(appDependencies)
	runServer := fmt.Sprintf("%s:%s", appDependencies.Config.HttpServer.ServerHost,
		appDependencies.Config.HttpServer.ServerPort)
	if errRun := appRoutes.Listen(runServer); errRun != nil {
		log.Fatal("run server error", errRun.Error())
	}
}
