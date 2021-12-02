package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	url := "https://www.im3db.com/chart/top/"

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".titleColumn").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a").Text()

		fmt.Println(i+1, title)
	})
}
