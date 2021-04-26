package main

import (
    "fmt"
    "os/exec"
    "io"
    "net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello")
}

func callBashToCheat() {
    //This website saved me, sorry it's hacky
    //https://zetcode.com/golang/exec-command/
	prg := "./app.sh"
	//arg1 := ndays
	//arg2 := symbol

	cmd := exec.Command(prg, ndays, symbol)
	stdout, err := cmd.Output()

	if err != nil {
	    fmt.Println(err.Error())
	    return
	}
	fmt.Print(string(stdout))
}

func getStocks(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    symbol := query.Get("symbol")   //should force this to uppercase
    ndays := query.Get("ndays")     //should force this to be int?
    w.WriteHeader(200)
    w.Header().Set("Inksworks", "gladly brought to you by")
    w.Header().Set("Content-Type", "text/html; charset=utf=8")
    //w.Write([]byte(strings.Join(filters, ",")))
    w.Write([]byte(symbol))
}

func main() {
    http.HandleFunc("/", echoString)
    http.HandleFunc("/stocks", getStocks)

    fmt.Println("Stock Server Listening on :3000...")
    println("Enter this in your browser:  http://localhost:3000/example?name=jenny&phone=867-5309")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
