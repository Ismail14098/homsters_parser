package resident_parser

import (
	"context"
	"github.com/Ismail14098/homsters_parser/common"
	"github.com/Ismail14098/homsters_parser/database/models"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Parse(resident common.ResponseEstateForBounds, logger *log.Logger,ctx *context.Context){
	idStr := strconv.Itoa(int(resident.ID))
	logger.Println("Saved residentRecord id = " + idStr)

	path := "plans/"+idStr
	os.Mkdir(path,0777)

	db := (*ctx).Value("db").(*gorm.DB)
	var developer models.Developer
	result := db.Where("name = ?", resident.DeveloperName).Find(&developer)
	if result.RowsAffected == 0 {
		developer.Name = resident.DeveloperName
		db.Save(&developer)
	}

	residentDB := models.Resident{
		Model:                                  gorm.Model{ID: resident.ID},
		Name:                                   resident.Name,
		DeveloperID:                            developer.ID,
		CityName:                               resident.CityName,
		DistrictName:                           resident.DistrictName,
		SubDistrictName:                        resident.SubDistrictName,
		PricePerSqM:                            resident.PricePerSqM,
		Currency:                               resident.Currency,
		MinSize:                                resident.MinSize,
		MaxSize:                                resident.MaxSize,
		MinRoomCount:                           resident.MinRoomCount,
		MaxRoomCount:                           resident.MaxRoomCount,
		FloorCount:                             resident.FloorCount,
		CommissioningYear:                      resident.CommissioningYear,
		CommissioningQuarter:                   resident.CommissioningQuarter,
		ConstructionStatusLocalizedDescription: resident.ConstructionStatusLocalizedDescription,
		PhoneNumber:                            resident.PhoneNumber,
		DirectPhone:                            resident.DirectPhone,
		ComplexUrl:                             resident.ComplexUrl,
		IsComplexHasActiveConstructionStatus:   resident.IsComplexHasActiveConstructionStatus,
	}

	db.Save(&residentDB)

	complexUrl := resident.ComplexUrl
	complexUrl+="/flatplans"

	response, err := http.Get(complexUrl)

	logger.Println("GET Request url = " + complexUrl)
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var parsedFlatPlans []models.Flatplan

	document.Find("a.b-flatplan-snippet.swiper-slide").Each(func(i int, selection *goquery.Selection) {
		//download image and place at folder
		img := selection.Find("img.b-flatplan-snippet__logo")
		url, _ := img.Attr("data-lazy-load-url")

		fileName := ""
		if url != "" {
			fileName = imageDownload(url, path, i, logger)
		}


		//parse data
		roomCountEl := selection.Find("h2")
		roomCountSlice := strings.Split(roomCountEl.Text(), "-")
		roomCountSlice = strings.Split(roomCountSlice[0]," ")
		roomCount, _ := strconv.Atoi(roomCountSlice[len(roomCountSlice)-1])

		sqMEl := selection.Find("span.b-flatplan-snippet__size")
		sqMStr := sqMEl.Text()
		sqMSlice := strings.Split(sqMStr, " m")
		sqMSlice = strings.Split(sqMSlice[0]," ")
		numStr := strings.Replace(sqMSlice[len(sqMSlice)-1], ",", ".", 1)
		sqMRaw, _ := strconv.ParseFloat(numStr,32)
		sqM := math.Round(sqMRaw*10)/10

		levelEl := selection.Find("span.b-flatplan-snippet__level")
		levelStr := levelEl.Text()
		levelSlice := strings.Split(levelStr,"-")
		var minLevel int
		var maxLevel int
		if len(levelSlice) != 1{
			levelMinSlice := strings.Split(levelSlice[0], " ")
			levelMaxSlice := strings.Split(levelSlice[1], " ")
			minLevel, _ = strconv.Atoi(levelMinSlice[len(levelMinSlice)-1])
			maxLevel, _ = strconv.Atoi(levelMaxSlice[0])
		} else {
			levelSlice = strings.Split(levelStr, " этаж")
			levelSlice = strings.Split(levelSlice[0]," ")
			minLevel, _ = strconv.Atoi(levelSlice[len(levelSlice)-1])
			maxLevel = minLevel
		}

		flatplan := models.Flatplan{
			ResidentID: resident.ID,
			RoomCount: uint(roomCount),
			SqM:       sqM,
			MinLevel:  uint(minLevel),
			MaxLevel:  uint(maxLevel),
			Image: fileName,
		}
		db.Save(&flatplan)

		parsedFlatPlans = append(parsedFlatPlans, flatplan)
	})

	// Hide in db flatplans that not listed in homsters.kz
	var dbFlatPlans []models.Flatplan
	db.Where("residentID = ?", resident.ID).Find(&dbFlatPlans)

	for _, v := range dbFlatPlans{
		for _, v1 := range parsedFlatPlans {
			if v.Image == v1.Image {
				break
			}
		}
		db.Delete(&v)
	}

	response.Body.Close()
}

func imageDownload(url string, path string, i int, logger *log.Logger) string{
	response, err := http.Get(url)

	defer response.Body.Close()
	if err != nil {
		logger.Println("Couldn't get image at", url, "err:", err)
	}

	fileSlice := strings.Split(url,"/")
	fileName := fileSlice[len(fileSlice)-1]
	//fileExt := strings.Split(fileName,".")
	//fileName = strconv.Itoa(i)+"."+fileExt[1]
	file, err := os.Create(filepath.Join(path,fileName))
	defer file.Close()
	io.Copy(file,response.Body)

	return fileName
}
