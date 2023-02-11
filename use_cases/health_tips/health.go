package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fruit struct {
	Genus      string         `json:"genus"`
	Name       string         `json:"name"`
	Id         int            `json:"id"`
	Family     string         `json:"family"`
	Order      string         `json:"order"`
	Nutritions map[string]int `json:"nutritions"`
}

func getFruits(c *gin.Context) {
	log.SetPrefix("Error getting fruits: ")
	log.SetFlags(0)

	var fruitJson []Fruit
	apiURL := "https://fruityvice.com/api/fruit/all"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	errr := json.Unmarshal(body, &fruitJson)
	if errr != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, fruitJson)
}

func main() {
	r := gin.Default()
	r.GET("/fruits", getFruits)
	r.Run("localhost:3000")
}
