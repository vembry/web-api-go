package main

import (
	"fmt"

	"github.com/vembry/web-api-template-go/repository"
	"github.com/vembry/web-api-template-go/service"
)

func main() {
	// fmt.Println("hur dur!")
	// r := gin.Default()

	// init repositories
	healthRepo := repository.InitRepository(repository.HealthRepoParam{})

	//init services
	healthServ := service.InitService(service.HealthServiceParam{
		IHealthRepo: &healthRepo,
	})

	fmt.Println("health check:", healthServ.Check())
	fmt.Println("health validate:", healthServ.Validate(2))

	// //init handlers
	// handler.InitHandler(r)

	// r.Run("localhost:3000")
}
