package main

import "os"
import "fmt"
import "time"
import "github.com/cavaliercoder/grab"

func main(){
	var key = os.Getenv("ALPHA_VANTAGE_API_KEY")
	var ticker = os.Args[1]
	var url = "https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=" + ticker + "&apikey=" + key + "&datatype=csv"


// TODO would be nice to move to separate file
resp, _ := grab.Get("", url)
fmt.Println("Response - ", resp.HTTPResponse)
defer resp.HTTPResponse.Body.Close()

/*
	Waiting for file to be downloaded:
*/
t := time.NewTicker(time.Second)
defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Printf("%.02f%% complete\n", resp.Progress())

		case <-resp.Done:
			if err := resp.Err(); err != nil {
				fmt.Println("ERROR!!! - ", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully downloaded to %s\n", resp.Filename)
			return
		}
	}
}