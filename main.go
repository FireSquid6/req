package main


import (
  "net/http"
)


type Request struct {
  Method string
  Url string
  Body string
}


func main() {

}


func Fetch() (*http.Response, error) {
  client := &http.Client{}

  req, err := http.NewRequest("GET", "http://localhost:8080", nil)

  if err != nil {
    return &http.Response{}, err
  }

  res, err := client.Do(req)

  if err != nil {
    return &http.Response{}, err
  }

  return res, nil
}
