// handlers.houses.go

package handlers

import (
  "net/http"
  "../repository"
  "strconv"

  "github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
  houseRepository := repository.HouseRepository{}

  houses := houseRepository.FetchAll()

  // Call the HTML method of the Context to render a template
  c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "index.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Home Page",
      "payload": houses,
    },
  )

}

func GetHouse(c *gin.Context){
  houseRepository := repository.HouseRepository{}

  houseId, err := strconv.Atoi(c.Param("house_id"));

  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  house, err := houseRepository.FetchById(houseId)

  if err != nil {
      // If the article is not found, abort with an error
      c.AbortWithError(http.StatusNotFound, err)
    return
  }

  c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the article.html template
    "house.html",
    // Pass the data that the page uses
    gin.H{
        "title":   house.Title,
        "payload": house,
    },
)
}