package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type Data struct { // TODO: sync with Ghost export format
    DB []Person `json:"db"`
}


func main() {
    file, _ := ioutil.ReadFile("config.json") // TODO: get from cli args
    var data Data
    json.Unmarshal(file, &data)
    for i, person := range data.DB {
        fileName := fmt.Sprintf("person%d.md", i)
        content := fmt.Sprintf("Name: %s\nAge: %d", person.Name, person.Age)
        ioutil.WriteFile(fileName, []byte(content), 0644)
    }

}