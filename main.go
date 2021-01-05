package main

import (
	"fmt"
	"github.com/Ismail14098/homsters_parser/city_parser"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main(){
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file,"LOG: ",log.Ldate|log.Ltime|log.Lshortfile)

	client := &http.Client{
		//Timeout: 2 * time.Second,
	}

	request, err := http.NewRequest("GET","https://homsters.kz", nil)
	if err != nil {
		log.Fatal(err)
	}

	//request.Header.Add()
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	document.Find("a.b-city-card.js-city-card").Each(func(i int, s *goquery.Selection) {
		//For each item found, get the band and title
		cityLink, _ := s.Attr("href")
		city_parser.Parse(client, cityLink, logger)
	})

}

