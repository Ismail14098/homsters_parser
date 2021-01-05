package city_parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Parse(client *http.Client, addr string, logger *log.Logger){
	logger.Printf("City address : %v", addr)
	request, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	document.Find("a.b-snippet__link.js-complex-ite.js-go-to-complex-pag.b-snippet__search.b-snippet__search--call-colored").Each(func(i int, selection *goquery.Selection) {

	})

}
