package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Ismail14098/homsters_parser/common"
	"github.com/Ismail14098/homsters_parser/parser/resident_parser"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Parse(client *http.Client, ctx context.Context, logger *log.Logger){
	os.Mkdir("plans",0777)
	urlAddr := "https://homsters.kz/estate/estateforbounds"
	for true {
		ctx = common.GetHeadersAndCookies(client,ctx, logger)
		page := 1
		for true {
			// getting list of residents
			pageStr := strconv.Itoa(page)
			logger.Println("Url " + urlAddr + ", page = " + pageStr)

			body, err := json.Marshal(common.JSON{
				"page":                        pageStr,
				"sort":                        "Featured",
				"sortdirection":               "Ascending",
				"TypeOfSearchComplexOrEstate": "1",
				"ConstructionStatus":          "4",
				"RoomsCount":                  "447",
				"Type":                        "127",
				"Heating":                     "2047",
				"Parking":                     "31",
				"ResidentialClass":            "15",
				"TypeOfHouse":                 "127",
				"IndoorFinish":                "31",
			})

			request, err := http.NewRequest("POST", urlAddr, bytes.NewBuffer(body))

			common.ModifyRequest(request, &ctx)

			response, err := client.Do(request)
			if err != nil {
				log.Fatal(err)
			}

			type responseStruct struct {
				Ads []common.ResponseEstateForBounds
			}
			var result responseStruct

			json.NewDecoder(response.Body).Decode(&result)

			response.Body.Close()
			client.CloseIdleConnections()

			if len(result.Ads) == 0 {
				break
			}

			for _, resident := range result.Ads {
				if resident.IsComplexHasActiveConstructionStatus {
					resident_parser.Parse(resident, logger, &ctx)
				}
			}
			page++
			time.Sleep(5 * time.Second)
			//Filter new records from old one
			//newRecords := filterNewRecords(result.Ads, ctx, logger)
			//if len(newRecords) > 0 {
			//	for _, resident := range newRecords {
			//		resident_parser.Parse(resident, client, logger, ctx)
			//		break
			//	}
			//}

			//if list have old records than next page is not needed to be parsed
			//if len(newRecords) != len(result.Ads){
			//	time.Sleep(24 * time.Hour)
			//} else {
			//	page++
			//}
		}
		time.Sleep(24 * time.Hour)
	}
}

//func filterNewRecords(records []models.Resident, ctx *context.Context, logger *log.Logger) (newRecords []models.Resident){
//	//var newRecords []models.Resident
//	rdb := (*ctx).Value("rdb").(*redis.Client)
//	for _, v := range records {
//		boolCmd := rdb.HExists(*ctx,"resident", strconv.Itoa(int(v.ID)))
//		if b,err := boolCmd.Result(); err == nil && !b { // maybe first record is not parsed
//			logger.Println("Added new record with id = " + strconv.Itoa(int(v.ID)))
//			newRecords = append(newRecords, v)
//		}
//	}
//	return
//}

//func trackParsedStatus(residentID uint, ctx *context.Context, logger *log.Logger){
//	rdb := (*ctx).Value("rdb").(*redis.Client)
//	rdb.HSet(*ctx, "resident", strconv.Itoa(int(residentID)))
//	db := (*ctx).Value("db").(*gorm.DB)
//
//
//}

