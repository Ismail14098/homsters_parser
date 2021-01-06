package main

import (
	"context"
	"github.com/Ismail14098/homsters_parser/parser"
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

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	ctx := context.Background()

	//fmt.Printf("%+v\n", header)
	cookies := response.Cookies()
	ctx = context.WithValue(ctx, "cookies", cookies)

	parser.Parse(client, logger, &ctx)
	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
}

