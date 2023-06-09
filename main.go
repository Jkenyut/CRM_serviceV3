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
	migrate "github.com/rubenv/sql-migrate"
	"time"
)

// func main
func main() {
	//main function
	_ = godotenv.Load()
	db := db2.GormMysql()
	//path := filepath.Join("crm_service.sql")
	//c, err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Println("error")
	//}
	//sql := string(c)
	//err = db.Exec(sql).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//migrate
	m := &migrate.FileMigrationSource{Dir: "."}
	sql, _ := db.DB()
	n, err := migrate.Exec(sql, "mysql", m, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	router := gin.New()

	//use cors
	router.Use(cors.Default())

	//use helmet
	router.Use(helmet.Default())

	//use rate-limit
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute * 1,
		Limit: 20,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: entity.ErrorHandler,
		KeyFunc:      entity.KeyFunc,
	})

	//call all route
	router.Use(mw)
	actorHandler := actor.NewRouter(db)
	actorHandler.Handle(router)

	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	errRouter := router.Run(":8081") // run gin
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
