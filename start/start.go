package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Card struct {
	Url, Image, Name, Percentage string
	Price                        []*cardPrices
}

type cardPrices struct {
	MainAnchor string `selector:"a"`
	Permalink  string `selector:".CardPrices_price__NBGhV span"`
	Comment    string
}

func main() {

	res, err := http.Get("https://edhrec.com/top")
	if err != nil {
		log.Fatal("Failed to connect to the target page", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	// convert the response buffer to bytes
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to fetch html doc", err)
	}

	doc.Find("div.Card_container__Ng56K").Each(func(i int, p *goquery.Selection) {
		card := Card{}
		card.Name = p.Find("span.Card_name__Mpa7S").Text()
		// cardImageInfo := p.Find("a")

		links := doc.Find("a").Map(func(i int, a *goquery.Selection) string {
			link, _ := a.Attr("href")
			return link
		})
		fmt.Println(links)
		// if exist {
		// 	fmt.Println(src)
		// 	card.Image = src
		// }
		// fmt.Println(card)
	})
}
