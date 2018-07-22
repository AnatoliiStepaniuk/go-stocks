package main

import "os"
import "fmt"
import "github.com/AnatoliiStepaniuk/go-stocks/http/utils"

func main(){
	var key = os.Getenv("ALPHA_VANTAGE_API_KEY")
	var ticker = os.Args[1]
	var url = "https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&outputsize=compact&symbol=" + ticker + "&apikey=" + key + "&datatype=csv"

	var fileName = utils.Download("", url)
	fmt.Printf("Received name of file downloaded to %s\n", fileName)
			
}
