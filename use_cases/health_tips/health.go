package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

var fruitJson []Fruit

func getFruits(c *gin.Context) {
	log.SetPrefix("Error getting fruits: ")
	log.SetFlags(0)

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

func extractFruit(c *gin.Context) {
	/*
		Structure of algorithm.

		--> retrieve list of fruits from the query string.
		--> Store list of fruits name in a []string
		--> make all first letter capitalized. in the list of fruits.
		--> loop through list of fruit names to find information about fruit and remove fruit if name not found returning a statusnotfound with custom error message.
		--> calculate total amount of calories from fruits list.
		--> return statusOk with total calories. in a custom message.
	*/

	// Retrieve the list of fruit parameters inputed by user and check if user input is found, else return statusbad request.
	fruits := strings.Split(c.Param("fruits"), ",")
	if len(fruits) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No fruit selected"})
		return
	}

	// capitalize first letter of each fruit name
	// and run a loop search in the fruitjson for march
	var fruitsCapitalized = make([]string, len(fruits))
	var fruitsCalories = make([]int, len(fruits))
	for i, v := range fruits {
		fruit := string(v)
		fruitsCapitalized[i] = strings.ToUpper(fruit[:1]) + fruit[1:]
		for _, j := range fruitJson {
			if j.Name == fruitsCapitalized[i] {
				fruitsCalories[i] = j.Nutritions["calories"]
			}
		}
	}

	if len(fruitsCalories) != len(fruitsCapitalized) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error Message:": "Invalid fruit names"})
	}
	// Calculate total calories
	// I will accomplish this by summing the calories in the fruitscalories list.
	totalCalories := 0
	for _, v := range fruitsCalories {
		totalCalories += v
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Total Calories: ": totalCalories})
}

func main() {
	r := gin.Default()
	r.GET("/fruits", getFruits)
	r.GET("/extractFruit/:fruits", extractFruit)
	r.Run("localhost:3000")
}
