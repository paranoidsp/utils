package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Init some globals
type Config struct {
    Token string               `json:"token"`
    Users map[string]string    `json:"users"`
}

var SecretConfig Config

func init () {
    configFile, err := ioutil.ReadFile("secrets/notify-config.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(configFile ,  &SecretConfig)
    if err != nil {
        panic(err)
    }
    log.Println("Loaded secrets from config")
}


func SendNotification (user, message string) (error) {

    // Send pushover notification
    data := url.Values{
        "message": {message},
        "user": {SecretConfig.Users[user]},
        "token": {SecretConfig.Token},
    }

    resp, err := http.PostForm("https://api.pushover.net/1/messages.json", data)
    // log.Println("Finished pushover request")
    if err != nil {
        panic(err)
    }

    // Read response
    respText, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        log.Fatal("Got non-200 response from pushover: ", resp.StatusCode)
    }

    // Unmarshal and check status
    var respJson map[string]interface{}
    err = json.Unmarshal(respText, &respJson)
    if err != nil {
        panic(err)
    }

    status, ok := respJson["status"]
    if !ok {
        panic(ok)
    }

    log.Println("Response: ", status)
    // Catch and throw err if exception
    return nil
}
