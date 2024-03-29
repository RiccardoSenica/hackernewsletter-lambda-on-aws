package hackernews

import (
	"encoding/json"
	"fmt"
	"hackernewsletter/db"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HackernewsTable() string {
	return "hacker_news_table"
}

func HackernewsTemplate() string {
	return "hackernews_template"
}

func GetTopNewsIds(url string) (response []string) {
	fmt.Println("Retrieving news...")

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyString := string(body)
	response = strings.Split(bodyString[1:len(bodyString)-1], ",")

	fmt.Println("Done.")

	return response
}

func GetNewsById(id string, url string) (response db.News) {
	fmt.Printf("Retrieving data for news with id %v...\n", id)

	news_url := strings.ReplaceAll(url, "{ID}", id)
	res, err := http.Get(news_url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &response)

	fmt.Printf("News with id %v retrieved.\n", id)

	return response
}
