package main

import (
	"net/http"
	"rest-api-golang-pemula/config"
	"rest-api-golang-pemula/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.InitDB()
	//DBConn := &controllers.DBConn{DB: db}

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// CRUD -  Create, Read, Update & Delete
	//router.POST("/create-person", DBConn.CreatePerson)

	routers.PersonRoute(router, db)

	router.Run(":8080")

}
