package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	getMe := GetMeT{}
	test := []byte(`{
		"ok": true,
		"result": {
			"id": 1397587106,
			"is_bot": true,
			"first_name": "TestIntensiveGo",
			"username": "Test123IntensiveGoBot",
			"can_join_groups": true,
			"can_read_all_group_messages": false,
			"supports_inline_queries": false
		}
	}`)

	err := json.Unmarshal(test, &getMe)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%v", getMe)
	fmt.Println(getMe.Result.Username)
}
