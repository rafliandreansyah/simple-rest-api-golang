package controllers

import (
	"net/http"
	"rest-api-golang-pemula/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) CreatePerson(c *gin.Context) {
	var person structs.Person

	age, err := strconv.Atoi(c.PostForm("age"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	person.FirstName = c.PostForm("first_name")
	person.LastName = c.PostForm("last_name")
	person.Age = age

	Conn.DB.Create(&person)

	result := gin.H{
		"result": person,
	}

	c.JSON(http.StatusCreated, result)
}

func (Conn *DBConn) GetPersonById(c *gin.Context) {
	var person structs.Person
	var result gin.H

	id := c.Param("id")

	err := Conn.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"message": "User not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	c.JSON(http.StatusOK, person)
}

func (Conn *DBConn) GetPersons(c *gin.Context) {

	var (
		persons []structs.Person
		result  gin.H
	)

	err := Conn.DB.Find(&persons).Error
	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	if len(persons) >= 0 {
		result = gin.H{
			"persons": persons,
		}
		c.JSON(http.StatusOK, result)
	}

}

func (Conn *DBConn) UpdatePerson(c *gin.Context) {

	var (
		person structs.Person
		result gin.H
		err    error
	)

	id := c.Query("id")

	err = Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"message": "User not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	age, err := strconv.Atoi(c.PostForm("age"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	person.FirstName = c.PostForm("first_name")
	person.LastName = c.PostForm("last_name")
	person.Age = age

	err = Conn.DB.Save(&person).Error
	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"message": "person updated",
		"person":  person,
	}
	c.JSON(http.StatusOK, result)

}

func (Conn *DBConn) DeletePerson(c *gin.Context) {

	var (
		person structs.Person
		err    error
		result gin.H
	)

	id := c.Query("id")

	err = Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"message": "User not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	err = Conn.DB.Delete(&person, id).Error
	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"message": "person deleted",
	}
	c.JSON(http.StatusOK, result)

}
