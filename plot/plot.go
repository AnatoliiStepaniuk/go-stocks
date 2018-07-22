package main


import (
	"os"
	"fmt"
	"github.com/AnatoliiStepaniuk/go-stocks/http/utils"
	"log"
	"io"
	"encoding/csv"
	"strconv"
	"time"
	"github.com/wcharczuk/go-chart"
)

func main() {
	var ticker = os.Args[1]
	var key = os.Getenv("ALPHA_VANTAGE_API_KEY")
	var url = "https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&outputsize=compact&symbol=" + ticker + "&apikey=" + key + "&datatype=csv"
	var baseFileName = "plot/files/monthly_" + ticker
	var intputFileName = baseFileName + ".csv"
	var outputFileName = baseFileName + ".png"

	/*
		Download file if not present
	*/
	f, _ := os.Open(intputFileName)
	
	if(f == nil){
		fmt.Println("File is absent, downloading")
		intputFileName = utils.Download("plot/files", url)
		f, _ = os.Open(intputFileName)
	}
    defer f.Close()
    
    /*
    	Read file contents
    */
    r := csv.NewReader(f)
    dates := make([]time.Time,0)
    values := make([]float64, 0)
	_, _ = r.Read()	 
		 for{
	    record, e := r.Read()
	    if(e == io.EOF){
	    	break
	    }
	    if(e != nil){
	    	log.Fatal(e)
	    }
	    date, _ := time.Parse(chart.DefaultDateFormat, record[0])
	    dates = append(dates, date)
	    fl, efl := strconv.ParseFloat(record[4], 64)
	    check(efl)
	    values = append(values, fl)
    }

    /*
		Drawing a chart
    */
	priceSeries := chart.TimeSeries{
		Name: ticker,
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: dates,
		YValues: values,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:        chart.Style{Show: true},
			TickPosition: chart.TickPositionBetweenTicks,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{Show: true},
		},
		Series: []chart.Series{
			priceSeries,
		},
	}
	outputFile, _ := os.Create(outputFileName)
	err := graph.Render(chart.PNG, outputFile)
	check(err)

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

