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

type SendMessageT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}

type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}
type GetUpdatesMessageT struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

type GetUpdatesResultT struct {
	UpdateId int                `json:"update_id"`
	Message  GetUpdatesMessageT `json:"message,omitempty"`
}

type MessageT struct {
	MessageId int                          `json:"message_id"`
	From      GetUpdatesResultMessageFromT `json:"from"`
	Chat      GetUpdatesResultMessageChatT `json:"chat"`
	Date      int                          `json:"date"`
	Text      bool                         `json:"text"`
}

type GetUpdatesResultMessageFromT struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type GetUpdatesResultMessageChatT struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

const telegramBaseUrl = "https://api.telegram.org/bot"
const telegramToken = "1397587106:AAF8socYPHdzdwbe6FuYW0M0GZicYlyGkLg"

const methodGetMe = "getMe"
const methodGetUpdates = "getUpdates"

func main() {

	//fmt.Println(getUrlByMethod(methodGetMe))

	body := getBodyByUrlAndData(getUrlByMethod(methodGetUpdates), []byte(""))

	//fmt.Printf("%s", body)

	getUpdates := GetUpdatesT{}
	err := json.Unmarshal(body, &getUpdates)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, item := range getUpdates.Result {
		fmt.Println(item.Message.Text)
	}

	//fmt.Printf("%v", getUpdates)
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
