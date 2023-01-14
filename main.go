package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct { // TODO: sync with Ghost export format
    Name string `json:"name"`
    Age  int    `json:"age"`
    Cars []struct {
        Brand string `json:"brand"`
        Model string `json:"model"`
        Year int    `json:"year"`
    }
}

func main() {
    file, _ := ioutil.ReadFile("config.json") // TODO: get from cli args
    var config Config
    json.Unmarshal(file, &config)
    fmt.Println(config.Name)
    fmt.Println(config.Age)
    fmt.Println(config.Cars[0].Brand)
    fmt.Println(config.Cars[0].Model)
    fmt.Println(config.Cars[0].Year)
    fmt.Println(config.Cars[1].Brand)
    fmt.Println(config.Cars[1].Model)
    fmt.Println(config.Cars[1].Year)
    
}