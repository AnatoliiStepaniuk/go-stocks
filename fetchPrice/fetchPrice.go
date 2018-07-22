package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "regexp"

func main() {

// Reading stock ticker from command line
var ticker = os.Args[1]

// Fetching data
var url = "https://finance.yahoo.com/quote/" + ticker + "/key-statistics?p=" + ticker
resp, err := http.Get(url)
defer resp.Body.Close()
if err != nil {
    fmt.Println("ERROR OCCURED!")
}
bytes, _ := ioutil.ReadAll(resp.Body)

// Parsing value (Trailing P/E)
var foundCoarse = regexp.MustCompile("Trailing P\\/E[\\S\\s]*\\d+\\.\\d{2}").FindString(string(bytes))
var pe = regexp.MustCompile("\\d+\\.\\d{2}").FindString(foundCoarse)

// Display
fmt.Println("Trailing P/E for ", ticker, " is ", pe)
}