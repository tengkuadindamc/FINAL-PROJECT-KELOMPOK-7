package main

import (
	"final-project3/routes"
	"final-project3/utils/postgres"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		HttpMainHandler()
		defer wg.Done()
	}()
	wg.Wait()
}

func HttpMainHandler() {
	g := gin.Default()
	db := postgres.NewConnection(postgres.BaseConfig()).Database

	routes.InitHttpRoute(g, db)

	g.Run()
}
