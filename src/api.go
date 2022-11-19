package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "log"
    "fmt"
)
 
type GithubRequest struct {
    link string // 'https://api.github.com/repos/trovalds/linux' or 'https://api.github.com/users/trovalds'
}


func NewRequest(link string) map[string]interface{} {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/%s", link),  nil)
    if err != nil {
        log.Fatalf("Error creating http request: %s", err.Error())
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Fatal(err)
    }

    var payload map[string]interface{}
    err = json.Unmarshal([]byte(body), &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
    return payload
}
