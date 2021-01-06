package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ismail14098/homsters_parser/common"
	"github.com/Ismail14098/homsters_parser/parser/resident_parser"
	"log"
	"net/http"
	"strconv"
)

func Parse(client *http.Client, logger *log.Logger, ctx *context.Context){
	//"IsComplexHasActiveConstructionStatus": true,
	page := 1
	urlAddr := "https://homsters.kz/estate/estateforbounds"
	for true {
		pageStr := strconv.Itoa(page)
		logger.Println("Url " + urlAddr + ", page = " + pageStr)

		body, err := json.Marshal(common.JSON{
				"page": pageStr,
				"sort": "Featured",
				"sortdirection": "Ascending",
				"TypeOfSearchComplexOrEstate": "1",
				"ConstructionStatus": "4",
				"RoomsCount": "447",
				"Type": "127",
				"Heating": "2047",
				"Parking": "31",
				"ResidentialClass": "15",
				"TypeOfHouse": "127",
				"IndoorFinish": "31",
		})

		request, err := http.NewRequest("POST",urlAddr, bytes.NewBuffer(body))

		modifyRequest(request, ctx)

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		var result common.JSON
		json.NewDecoder(response.Body).Decode(&result)
		fmt.Printf("%+v\n", result)

		ads := result["ads"].([]map[string]interface{})
		for _, resident := range ads{
			resident_parser.Parse(&resident, logger)
		}

		response.Body.Close()
		break
	}
}

func modifyRequest(request *http.Request, ctx *context.Context){
	request.Header.Set("Content-Type","application/json")
	request.Header.Set("Content-Encoding", "gzip, deflate, br")
	request.Header.Set("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Mobile Safari/537.36")
	request.Header.Set("Connection", "keep-alive")

	cookies := (*ctx).Value("cookies").([]*http.Cookie)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}
}
