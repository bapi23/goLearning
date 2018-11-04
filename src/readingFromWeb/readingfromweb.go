package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "strings"
 )

type Tour struct {
  Name, Price, PackageTitle string
}

func toursFromJSON(content string) []Tour {
  tours := make([]Tour, 0, 20)

  decoder := json.NewDecoder(strings.NewReader(content))
  decoder.Token()
  //checkError(err)
  var tour Tour
  for decoder.More() {
    decoder.Decode(&tour)
    //checkError(err)
    tours = append(tours, tour)
  }
  return tours
}

func main() {
  url := "http://services.explorecalifornia.org/json/tours.php"

  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Response type %T\n", resp)
  defer resp.Body.Close()

  bytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  content := string(bytes)
  fmt.Print(content)

  tours := toursFromJSON(content)
  fmt.Println(tours)
}
