package main

import (
	"context"
	"github.com/Ismail14098/homsters_parser/database"
	"github.com/Ismail14098/homsters_parser/parser"
	"github.com/Ismail14098/homsters_parser/redis"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main(){
	// load .env
	err := godotenv.Load()

	// create context
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	// create logger
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file,"LOG: ",log.Ldate|log.Ltime|log.Lshortfile)

	// create connection to postgresql
	db := database.Initialize(logger)
	ctx = context.WithValue(ctx, "db", db)

	// create connection to redis
	rdb := redis.Initialize()
	ctx = context.WithValue(ctx, "rdb", rdb)

	client := &http.Client{
		//Timeout: 2 * time.Second,
	}

	// Retrieve cookies from host
	request, err := http.NewRequest("GET","https://homsters.kz", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	cookies := response.Cookies()
	ctx = context.WithValue(ctx, "cookies", cookies)

	parser.Parse(client, &ctx, logger)
}

