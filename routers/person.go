package routers

import (
	"rest-api-golang-pemula/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PersonRoute(route *gin.Engine, db *gorm.DB) {

	DBConn := &controllers.DBConn{DB: db}

	// http://localhost:8080/1 -> GET with parameter
	route.GET("/:id", DBConn.GetPersonById)

	// http://localhost:8080/ -> GET
	route.GET("/", DBConn.GetPersons)

	// http://localhost:8080/create-person -> POST
	route.POST("/create-person", DBConn.CreatePerson)

	// http://localhost:8080/update-person?id=1
	route.PUT("/update-person", DBConn.UpdatePerson)

	// http://localhost:8080/delete-person?id=1
	route.DELETE("/delete-person", DBConn.DeletePerson)
}
