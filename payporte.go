package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {

	file, err := os.Create("porte001.csv")
	if err != nil {
		log.Fatalf("Could not create file: %q\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("payporte.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML(".product-item", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".product-item-info > .product-item-details > h2 > a"),
			e.ChildText(".product-item-info > .product-item-details > .price-box > .normal-price > .price-container > .price-wrapper > .price"),
		})
	})

	for i := 0; i < 10; i++ {
		fmt.Printf("Scraping Page: %d\n", i)

		c.Visit("https://payporte.com/clearance-sales.html?p=" + strconv.Itoa(i))
	}

	log.Println(c)
}
