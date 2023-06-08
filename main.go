package main

import (
	"crm_serviceV3/entity"
	"crm_serviceV3/modules/actor"
	"crm_serviceV3/modules/customer"
	db2 "crm_serviceV3/utils/db"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	//main function
	_ = godotenv.Load()
	db := db2.GormMysql()
	router := gin.New()

	router.Use(cors.Default())
	router.Use(helmet.Default())
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute * 1,
		Limit: 20,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: entity.ErrorHandler,
		KeyFunc:      entity.KeyFunc,
	})
	router.Use(mw)
	actorHandler := actor.NewRouter(db)
	actorHandler.Handle(router)

	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
