package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)
func main() {
    var jsonStr =`
{
  "name": "山田太郎",
  "age": 30,
  "isStudent": false,
  "address": {
    "street": "桜通り1-2-3",
    "city": "東京",
    "zip": "100-0001"
  },
  "hobbies": ["読書", "映画鑑賞", "旅行"]
}`

    name, err := jsonparser.GetString([]byte(jsonStr), "name")
    if err != nil {
        return
    }
    fmt.Println(name)

    age, err := jsonparser.GetInt([]byte(jsonStr), "age")
    if err != nil {
        return
    }
    fmt.Println(age)

    _, err = jsonparser.ArrayEach([]byte(jsonStr), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
        fmt.Println(string(value))
    }, "hobbies")
    if err != nil {
        return
    }
}