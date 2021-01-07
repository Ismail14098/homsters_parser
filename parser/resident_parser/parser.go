package resident_parser

import (
	"context"
	"fmt"
	"github.com/Ismail14098/homsters_parser/database/models"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Parse(resident models.Resident, client *http.Client, logger *log.Logger,ctx *context.Context){
	idStr := strconv.Itoa(int(resident.ID))
	logger.Println("Saved residentRecord id = " + idStr)

	path := "plans/"+idStr
	err := os.Mkdir(path,0777)
	if err != nil {
		return
	}

	db := (*ctx).Value("db").(*gorm.DB)
	res := db.Save(&resident)
	fmt.Println(res.Error)

	complexUrl := resident.ComplexUrl
	complexUrl+="/flatplans"

	request, err := http.NewRequest("GET", complexUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//common.ModifyRequest(request, ctx)
	request.Header.Set("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Mobile Safari/537.36")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept","*/*")
	request.Header.Set("Accept-Encoding","gzip, deflate, br")
	cookies := (*ctx).Value("cookies").([]*http.Cookie)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	client1 := &http.Client{}

	response, err := client1.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	logger.Println("GET Request url = " + complexUrl)
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result")
	v, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(v))

	document.Find("img.b-flatplan-snippet__logo").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("src")
		imageDownload(url, path, logger)
	})

	response.Body.Close()
}

func imageDownload(url string, path string, logger *log.Logger){
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		logger.Println("Couldn't get image at", url, "err:", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Println("Couldn't get image at", url, "err:", err)
		return
	}

	err = ioutil.WriteFile(path, data, 0755)
	if err != nil {
		logger.Println("Couldn't write image")
	}
}
