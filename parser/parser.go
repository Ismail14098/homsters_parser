package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ismail14098/homsters_parser/common"
	"github.com/Ismail14098/homsters_parser/database/models"
	"github.com/Ismail14098/homsters_parser/parser/resident_parser"
	"log"
	"net/http"
	"strconv"
)

func Parse(client *http.Client, ctx *context.Context, logger *log.Logger){
	//"IsComplexHasActiveConstructionStatus": true,
	page := 1
	urlAddr := "https://homsters.kz/estate/estateforbounds"
	for true {
		// getting list of residents
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

		request, err := http.NewRequest("POST", urlAddr, bytes.NewBuffer(body))

		common.ModifyRequest(request, ctx)

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		type responseStruct struct {
			Ads []models.Resident
		}
		var result responseStruct

		err = json.NewDecoder(response.Body).Decode(&result)

		//Filter new records from old one
		//newRecords := filterNewRecords(result.Ads,ctx,logger)
		//if len(newRecords) > 0 {
		//	for _, resident := range newRecords {
		//		resident_parser.Parse(resident, client, logger, ctx)
		//	}
		//}

		for _, resident := range result.Ads {
			resident_parser.Parse(resident, client, logger, ctx)
			break
		}

		//if list have old records than next page is not needed to be parsed
		//if len(newRecords) != len(result.Ads){
		//	time.Sleep(24 * time.Hour)
		//} else {
		//	page++
		//}

		response.Body.Close()
		//Test
		break
	}
}

func filterNewRecords(records []models.Resident, ctx *context.Context, logger *log.Logger) (newRecords []models.Resident){
	//var newRecords []models.Resident
	//for _, v := range records {
	//	rdb := (*ctx).Value("rdb").(*redis.Client)
	//	boolCmd := rdb.HExists(*ctx,"resident",strconv.Itoa(int(v.ID)))
	//	if b,err := boolCmd.Result(); err != nil && !b { // maybe first record is not parsed
	//		logger.Println("Added new record with id = " + strconv.Itoa(int(v.ID)))
	//		newRecords = append(newRecords, v)
	//	}
	//}
	return
}

