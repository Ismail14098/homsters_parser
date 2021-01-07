package common

import (
	"context"
	"net/http"
)

type JSON map[string]interface{}

func ModifyRequest(request *http.Request, ctx *context.Context){
	request.Header.Set("Content-Type","application/json")
	request.Header.Set("Content-Encoding", "gzip, deflate, br")
	request.Header.Set("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Mobile Safari/537.36")
	request.Header.Set("Connection", "keep-alive")

	cookies := (*ctx).Value("cookies").([]*http.Cookie)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}
}
