package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../data"
	"fmt"
)

// Handler for HTTP Get - "/showtimes"
// Returns all Showtime documents
func GetShowTimes(c *gin.Context) {
	// Get all showtimes form repository
	showtimes := data.GetAll()
	c.JSON(http.StatusOK,GoTestResource{Data: showtimes})
}

func GetShowTimeById(c *gin.Context) {
	var err error
	id  := c.Param("id")
	// Create new context
	fmt.Println(id)
	showtimes,err := data.GetAllbyID(id)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,GoTestResourceOne{Data: showtimes})
}

func Update(c *gin.Context) {
	var dataResource GoTestResourceOne
	c.Bind(&dataResource.Data)
	err := data.Update(dataResource.Data)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, dataResource)
}
