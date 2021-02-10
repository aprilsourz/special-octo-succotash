package main

import (
    "github.com/arlo-feirman.special-octo-succotash/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r.GET("/animal/:name", func(c *gin.Context) {
        animal, err := database.GetAnimal(c.Param("name"))
        if err != nil {
          c.String(404, err.Error())
          return
        }
        c.JSON(200, animal)
      })
}


