package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{"Alice", 30}
    jsonString, _ := json.Marshal(person)
    fmt.Println(string(jsonString))

    jsonText := `{"Name":"Bob","Age":25}`
    var person2 Person
    json.Unmarshal([]byte(jsonText), &person2)
    fmt.Println(person2)
}
