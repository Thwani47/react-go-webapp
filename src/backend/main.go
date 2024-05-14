package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Thwani47/react-go-webapp/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Workding directory: ", wd)

	db, err := sql.Open("sqlite", wd+"/database/database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()

		if err != nil {
			log.Fatal(err)
		}
	}(db)

	router := gin.Default()

	router.POST("/users", func(c *gin.Context) { handlers.CreateUser(c, db) })
	router.POST("/channels", func(c *gin.Context) { handlers.CreateChannel(c, db) })
	router.POST("/messages", func(c *gin.Context) { handlers.CreateMessage(c, db) })

	router.GET("/channels", func(c *gin.Context) {
		log.Println("endpoint called")
		handlers.ListChannels(c, db)
	})
	router.GET("/messages", func(c *gin.Context) { handlers.ListMessages(c, db) })
	router.POST("/login", func(c *gin.Context) { handlers.Login(c, db) })

	err = router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
