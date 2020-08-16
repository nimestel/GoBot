package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}

type GetMeResultT struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

const telegramBaseUrl = "https://api.telegram.org/bot"
const telegramToken = "1397587106:AAF8socYPHdzdwbe6FuYW0M0GZicYlyGkLg"

const methodGetMe = "getMe"

func main() {

	fmt.Println(getUrlByMethod(methodGetMe))

	body := getBodyByUrlAndData(getUrlByMethod(methodGetMe), []byte(""))
	//fmt.Printf("%s", body)
	getMe := GetMeT{}
	err := json.Unmarshal(body, &getMe)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%v", getMe)
}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + telegramToken + "/" + methodName
}

func getBodyByUrlAndData(url string, data []byte) []byte {
	r := bytes.NewReader(data)
	response, err := http.Post(url, "application/json", r)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}
