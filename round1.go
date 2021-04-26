package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
  Day string

}

func main() {
	birdJson := `{
	"Meta Data":
		{"name": "carl"},
	"Time Series (Daily)": {
        "2021-04-21":
		{"1. 1": "111"}
	},
        "2021-04-22":
		{"1. 2": "222"}
	},
        "2021-04-23":
		{"1. 3": "333"}
	}
	}`


	var f map[string]interface{}

	json.Unmarshal([]byte(birdJson), &f)

	fmt.Println(f["Meta Data"])

	fmt.Println(f["Time Series (Daily)"])

	var count = 0
	for key, value := range f["Time Series (Daily)"].(map[string]interface{})   {
		fmt.Println(value)
		fmt.Println(key)

		count++
	}

	fmt.Println(count)


}
