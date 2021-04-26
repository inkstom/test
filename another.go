package main

import "time"
import "net/http"
import "encoding/json"

//n, err := decimal.NewFromString("-123.4567")
//n.String() // output: "-123.4567"

//Don't use the default HTTP client, always specify the timeout in http.Client according to your use case
var httpClient = &http.Client{
  Timeout: time.Second * 3,
}

func getJson(url string, target interface{}) error {
    r, err := httpClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func
blogTitles, err := GetLatestBlogTitles("https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT")
if err != nil {
    log.Println(err)
}
fmt.Println("Blog Titles:")
fmt.Printf(blogTitles)
}

type Day_info struct {
    Close               string		`json:"4. close"`
}

func main() {
    getClose := Day_info{}
    getJson("https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT", &getClose)
    println(getJson)
    println(getClose.Close)
}
