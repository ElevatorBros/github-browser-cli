package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "encoding/json"
)

type Repo struct {
    Name string
    Description string
    Bar string
}

func (repo Repo) GetReadMe() string {
    url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/README.md", repo.Name, repo.GetDefaultBranch()) 
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error curling readme: %s", err.Error())
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil { 
        log.Fatalf("Error reading read me: %s", err.Error())
    }

    return string(body)
}


func (repo Repo) GetDefaultBranch() string {
    url := fmt.Sprintf("https://api.github.com/repos/%s", repo.Name)
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error curling api in GetDefaultBranch: %s", err.Error())
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil { 
        log.Fatalf("Error reading api: %s", err.Error())
    }

    var payload map[string]interface{}
    err = json.Unmarshal(body, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

    return fmt.Sprintf("%s", payload["default_branch"])
}


func lower(s string) string {
  // Convert it to ascii
  ascii := []byte(s)
  res := []byte(s)

  for i, j:= range ascii {
    // If uppercase
    if j >= 65 && j <= 90 {
      res[i] = j + 32
    } else {
      res[i] = j
    }
  }

  return string(res)
}
