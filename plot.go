package main


import (
"os"
"fmt"
"net/http"
"io/ioutil"
"regexp"
"github.com/cavaliercoder/grab"
"time"
)

func main() {

/*
	Reading stock ticker and time range from command line and downloading file
*/
if(len(os.Args) != 4){
	panic("Expecting 3 arguments - ticker timeFrom timeTo")
}

downloadFile(os.Args[1], os.Args[2], os.Args[3])




}

func downloadFile(ticker string, timeFrom string, timeTo string){

/*
	Getting crumb and cookie for downloading data:
*/
var url = "https://finance.yahoo.com/quote/" + ticker + "/key-statistics?p=" + ticker
resp, err := http.Get(url)
defer resp.Body.Close()
if err != nil {
    fmt.Println("ERROR OCCURED!")
}
bytes, _ := ioutil.ReadAll(resp.Body)
var crumb = regexp.MustCompile(`"crumb":"(\w{10,12})"`).FindAllStringSubmatch(string(bytes), -1)[0][1]
var cookie = resp.Cookies()[0]

/*
/*
	Using Grab library for downloading file (using crumb and cookie obtained earlier)
*/
var downloadUrl = "https://query1.finance.yahoo.com/v7/finance/download/" + ticker + "?period1=" + timeFrom + "&period2=" + timeTo + "&interval=1d&events=history&crumb="+crumb

var grabRequest, _ = grab.NewRequest("", downloadUrl)
grabRequest.HTTPRequest.AddCookie(cookie)
grabResp := grab.NewClient().Do(grabRequest)
defer grabResp.HTTPResponse.Body.Close()

/*
	Waiting for file to be downloaded:
*/
t := time.NewTicker(time.Second)
defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Printf("%.02f%% complete\n", grabResp.Progress())

		case <-grabResp.Done:
			if grabError := grabResp.Err(); grabError != nil {
				fmt.Println("ERROR!!! - ", grabError)
				os.Exit(1)
			}
			fmt.Printf("Successfully downloaded to %s\n", grabResp.Filename)
			return
		}
	}
}