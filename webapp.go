package main

import (
    "fmt"       //Basic I/O functions
    "log"       //Logging
    "net/http"  //Package http provides HTTP client and server implementations.
    "strconv"   //Converts integer to ascii
    "sync"      //Package sync provides basic synchronization primitives such as mutual exclusion locks
    //"encoding/json"
)

var counter int
// A Mutex is a mutual exclusion lock. A Mutex is used to provide a locking mechanism to ensure
// that only one Goroutine is running the critical section of code at any point in time to
// prevent race conditions from happening.

var mutex = &sync.Mutex{}
var apikey= "apikey=C227WD9W3LUVKVV9"

func echoString(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    // start goroutines to simulate writes, using the same pattern we did for reads.
    mutex.Lock()
    counter++
    fmt.Fprint(w, strconv.Itoa(counter))
    mutex.Unlock()

    fmt.Println ("Increment data:")
    fmt.Println (" name: " +r.URL.Query().Get("name"))
    fmt.Println (" truck: " +r.URL.Query().Get("truck"))
}

func getStocks(w http.ResponseWriter, r *http.Request) {

    resp, err := http.Get("https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    //https://lets-go.alexedwards.net/sample/02.06-url-query-strings.html
    // Extract the value of the id parameter from the query string and try to
    // convert it to an integer using the strconv.Atoi() function. If it can't
    // be converted to an integer, or the value is less than 1, we return a 404 page
    // not found response.

    // Keep going
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(resp))

}

func main() {
    http.HandleFunc("/", echoString)
    http.HandleFunc("/incrememnt", incrementCounter)
    http.HandleFunc("/stocks", getStocks)
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hi")
    })

    fmt.Println("Stock Server Listening on :3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
