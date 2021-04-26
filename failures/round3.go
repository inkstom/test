package main

import (
   "fmt"
   "encoding/json"
   "log"
   "net/http"

)

// Day specific data is exported, it models the data we receive.
type Day_info []struct {
    Close        string		`json:"4. close"`
}

//TextOutput is exported,it formats the data to plain text.
func (c Day_info) TextOutput() string {
p := fmt.Sprintf(
  "Close: %s\nSymbol : nAverage :",
  c[0].Close,)
   return p
}

// FetchCrypto will retrieve a url and return the body of the page.
func FetchCrypto(fiat string , crypto string) (string, error) {
//Build The URL string
    URL := "https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol="+crypto
    //URL := "https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT"
    fmt.Println(fiat)

    // If Get encounters any errors, it will exit 1.
    resp, err := http.Get(URL)
    if err != nil {
       log.Fatal("ooopsss an error occurred, please try again1")
    }
    defer resp.Body.Close()
 //Create a variable of the same type as our model
    var cResp Day_info
 //Decode the data
    if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
       log.Fatal("ooopsss! an error occurred, please try again2")
    }
 //Invoke the text output function & return it with nil as the error value
    return cResp.TextOutput(), nil
 }

func main() {
    //apikey :="demo"
    //symbol :="MSFT"

    crypto, err := FetchCrypto("printme", "MSFT")
    if err != nil {
        log.Println(err)
      }

  fmt.Println(crypto)
}
