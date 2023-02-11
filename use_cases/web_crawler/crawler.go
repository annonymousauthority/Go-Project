package crawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func CrawlPlatforms() {
	log.SetPrefix("Error finding Page: ")
	log.SetFlags(0)
	url := "https://fruityvice.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the HTML of the web page
	data, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(data)
	fmt.Println(resp)
}
