package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func api(method, uri, token string, data url.Values) ([]byte, error) {
	req, _ := http.NewRequest(method, "https://api.chatwork.com"+uri, strings.NewReader(data.Encode()))
	req.Header.Set("X-ChatWorkToken", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll(res.Body)
	return byteArray, err
}

type accountProfile struct {
	AccountId int    `json:"account_id"`
	Name      string `json:"name"`
	RoomId    int    `json:"room_id"`
}

func me(token string) (accountProfile, error) {
	res_json, err := api("GET", "/v1/me", token, nil)
	var profile accountProfile
	json.Unmarshal(res_json, &profile)
	return profile, err
}

type messageResponse struct {
	MessageId int `json:"message_id"`
}

func send_message(token string, room_id int, message string) (messageResponse, error) {
	data := url.Values{"body": {message}}
	res_json, err := api("POST", "/v1/rooms/"+strconv.Itoa(room_id)+"/messages", token, data)
	var msg_res messageResponse
	json.Unmarshal(res_json, &msg_res)
	return msg_res, err
}

func main() {
	token := os.Getenv("CHATWORK_TOKEN")

	profile, err := me(token)
	if err != nil {
		fmt.Println("invalid token")
		os.Exit(0)
	}

	var account_id int
	var room_id int

	flag.IntVar(&account_id, "to", profile.AccountId, "Account ID")
	flag.IntVar(&account_id, "t", profile.AccountId, "Account ID (short)")

	flag.IntVar(&room_id, "room", profile.RoomId, "Room ID")
	flag.IntVar(&room_id, "r", profile.RoomId, "Room ID (short)")
	flag.Parse()

	append_message := os.Getenv("CHATWORK_APPEND_MESSAGE")
	prepend_message := os.Getenv("CHATWORK_PREPEND_MESSAGE")
	var messages []string
	messages = flag.Args()

	if len(messages) > 0 {
		for _, message := range messages {
			_, err := send_message(token, room_id, prepend_message+message+append_message)
			if err != nil {
				fmt.Println("faild")
			}
		}
	} else {
		message, _ := ioutil.ReadAll(os.Stdin)
		_, err := send_message(token, room_id, prepend_message+string(message)+append_message)
		if err != nil {
			fmt.Println("faild")
		}
	}
}
