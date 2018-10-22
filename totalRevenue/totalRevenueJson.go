package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type kv struct {
	Key   string
	Value float64
}

const TAX = 0.15
const YEAR_HOURS = 8760.0

func main() {
	//aristocrates := []string {"EMR", "GPC", "PG", "MMM", "CINF", "KO", "JNJ", "LOW", "CL", "FRT", "HRL", "SWK", "TGT", "SYY", "BDX", "LEG", "PPG", "GWW", "KMB", "NUE", "PEP", "ABBV", "ABT", "SPGI", "WMT", "ADP", "ED", "ITW", "ADM", "WBA", "MCD", "PNR", "CLX", "MDT", "SHW", "BEN", "AFL", "XOM", "T", "CTAS", "MKC", "CAH", "TROW", "CVX", "ECL", "GD"}
	s_p_500 := []string{"ABT ", "ABBV ", "ACN ", "ACE ", "ADBE ", "ADT ", "AAP ", "AES ", "AET ", "AFL ", "AMG ", "A ", "GAS ", "APD ", "ARG ", "AKAM ", "AA ", "AGN ", "ALXN ", "ALLE ", "ADS ", "ALL ", "ALTR ", "MO ", "AMZN ", "AEE ", "AAL ", "AEP ", "AXP ", "AIG ", "AMT ", "AMP ", "ABC ", "AME ", "AMGN ", "APH ", "APC ", "ADI ", "AON ", "APA ", "AIV ", "AMAT ", "ADM ", "AIZ ", "T ", "ADSK ", "ADP ", "AN ", "AZO ", "AVGO ", "AVB ", "AVY ", "BHI ", "BLL ", "BAC ", "BK ", "BCR ", "BXLT ", "BAX ", "BBT ", "BDX ", "BBBY ", "BRK-B ", "BBY ", "BLX ", "HRB ", "BA ", "BWA ", "BXP ", "BSK ", "BMY ", "BRCM ", "BF-B ", "CHRW ", "CA ", "CVC ", "COG ", "CAM ", "CPB ", "COF ", "CAH ", "HSIC ", "KMX ", "CCL ", "CAT ", "CBG ", "CBS ", "CELG ", "CNP ", "CTL ", "CERN ", "CF ", "SCHW ", "CHK ", "CVX ", "CMG ", "CB ", "CI ", "XEC ", "CINF ", "CTAS ", "CSCO ", "C ", "CTXS ", "CLX ", "CME ", "CMS ", "COH ", "KO ", "CCE ", "CTSH ", "CL ", "CMCSA ", "CMA ", "CSC ", "CAG ", "COP ", "CNX ", "ED ", "STZ ", "GLW ", "COST ", "CCI ", "CSX ", "CMI ", "CVS ", "DHI ", "DHR ", "DRI ", "DVA ", "DE ", "DLPH ", "DAL ", "XRAY ", "DVN ", "DO ", "DTV ", "DFS ", "DISCA ", "DISCK ", "DG ", "DLTR ", "D ", "DOV ", "DOW ", "DPS ", "DTE ", "DD ", "DUK ", "DNB ", "ETFC ", "EMN ", "ETN ", "EBAY ", "ECL ", "EIX ", "EW ", "EA ", "EMC ", "EMR ", "ENDP ", "ESV ", "ETR ", "EOG ", "EQT ", "EFX ", "EQIX ", "EQR ", "ESS ", "EL ", "ES ", "EXC ", "EXPE ", "EXPD ", "ESRX ", "XOM ", "FFIV ", "FB ", "FAST ", "FDX ", "FIS ", "FITB ", "FSLR ", "FE ", "FSIV ", "FLIR ", "FLS ", "FLR ", "FMC ", "FTI ", "F ", "FOSL ", "BEN ", "FCX ", "FTR ", "GME ", "GPS ", "GRMN ", "GD ", "GE ", "GGP ", "GIS ", "GM ", "GPC ", "GNW ", "GILD ", "GS ", "GT ", "GOOGL ", "GOOG ", "GWW ", "HAL ", "HBI ", "HOG ", "HAR ", "HRS ", "HIG ", "HAS ", "HCA ", "HCP ", "HCN ", "HP ", "HES ", "HPQ ", "HD ", "HON ", "HRL ", "HSP ", "HST ", "HCBK ", "HUM ", "HBAN ", "ITW ", "IR ", "INTC ", "ICE ", "IBM ", "IP ", "IPG ", "IFF ", "INTU ", "ISRG ", "IVZ ", "IRM ", "JEC ", "JBHT ", "JNJ ", "JCI ", "JOY ", "JPM ", "JNPR ", "KSU ", "K ", "KEY ", "GMCR ", "KMB ", "KIM ", "KMI ", "KLAC ", "KSS ", "KRFT ", "KR ", "LB ", "LLL ", "LH ", "LRCX ", "LM ", "LEG ", "LEN ", "LVLT ", "LUK ", "LLY ", "LNC ", "LLTC ", "LMT ", "L ", "LOW ", "LYB ", "MTB ", "MAC ", "M ", "MNK ", "MRO ", "MPC ", "MAR ", "MMC ", "MLM ", "MAS ", "MA ", "MAT ", "MKC ", "MCD ", "MHFI ", "MCK ", "MJN ", "MMV ", "MDT ", "MRK ", "MET ", "KORS ", "MCHP ", "MU ", "MSFT ", "MHK ", "TAP ", "MDLZ ", "MON ", "MNST ", "MCO ", "MS ", "MOS ", "MSI ", "MUR ", "MYL ", "NDAQ ", "NOV ", "NAVI ", "NTAP ", "NFLX ", "NWL ", "NFX ", "NEM ", "NWSA ", "NEE ", "NLSN ", "NKE ", "NI ", "NE ", "NBL ", "JWN ", "NSC ", "NTRS ", "NOC ", "NRG ", "NUE ", "NVDA ", "ORLY ", "OXY ", "OMC ", "OKE ", "ORCL ", "OI ", "PCAR ", "PLL ", "PH ", "PDCO ", "PAYX ", "PNR ", "PBCT ", "POM ", "PEP ", "PKI ", "PRGO ", "PFE ", "PCG ", "PM ", "PSX ", "PNW ", "PXD ", "PBI ", "PCL ", "PNC ", "RL ", "PPG ", "PPL ", "PX ", "PCP ", "PCLN ", "PFG ", "PG ", "PGR ", "PLD ", "PRU ", "PEG ", "PSA ", "PHM ", "PVH ", "QRVO ", "PWR ", "QCOM ", "DGX ", "RRC ", "RTN ", "O ", "RHT ", "REGN ", "RF ", "RSG ", "RAI ", "RHI ", "ROK ", "COL ", "ROP ", "ROST ", "RLC ", "R ", "CRM ", "SNDK ", "SCG ", "SLB ", "SNI ", "STX ", "SEE ", "SRE ", "SHW ", "SIAL ", "SPG ", "SWKS ", "SLG ", "SJM ", "SNA ", "SO ", "LUV ", "SWN ", "SE ", "STJ ", "SWK ", "SPLS ", "SBUX ", "HOT ", "STT ", "SRCL ", "SYK ", "STI ", "SYMC ", "SYY ", "TROW ", "TGT ", "TEL ", "TE ", "TGNA ", "THC ", "TDC ", "TSO ", "TXN ", "TXT ", "HSY ", "TRV ", "TMO ", "TIF ", "TWX ", "TWC ", "TJK ", "TMK ", "TSS ", "TSCO ", "RIG ", "TRIP ", "FOXA ", "TSN ", "TYC ", "UA ", "UNP ", "UNH ", "UPS ", "URI ", "UTX ", "UHS ", "UNM ", "URBN ", "VFC ", "VLO ", "VAR ", "VTR ", "VRSN ", "VZ ", "VRTX ", "VIAB ", "V ", "VNO ", "VMC ", "WMT ", "WBA ", "DIS ", "WM ", "WAT ", "ANTM ", "WFC ", "WDC ", "WU ", "WY ", "WHR ", "WFM ", "WMB ", "WEC ", "WYN ", "WYNN ", "XEL ", "XRX ", "XLNX ", "XL ", "XYL ", "YHOO ", "YUM ", "ZBH ", "ZION ", "ZTS "}

	//fmt.Println("===== Dividend Aristocrates =====")
	//roiMap1 := getRoiMap(aristocrates)
	//fmt.Println("Average - ", fmt.Sprintf("%.2f", getAverage(getValues(roiMap1))))
	//printRoiMap(roiMap1)

	fmt.Println("===== S&P 500 =====")
	roiMap2 := getRoiMap(s_p_500)
	fmt.Println("Average - ", fmt.Sprintf("%.2f", getAverage(getValues(roiMap2))))
	printRoiMap(roiMap2)

}

func getValues(m map[string]float64) []float64 {
	sl := make([]float64, 0)
	for _, v := range m {
		sl = append(sl, v)
	}
	return sl
}

func getAverage(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func printRoiMap(roiMap map[string]float64) {
	sl := make([]kv, 0)
	for k, v := range roiMap {
		sl = append(sl, kv{k, v})
	}

	sort.Slice(sl, func(i, j int) bool { return sl[i].Value > sl[j].Value })

	for _, e := range sl {
		printRoi(e.Key, e.Value)
	}
}

func getRoiMap(tickers []string) map[string]float64 {

	roiMap := make(map[string]float64)

	for _, ticker := range tickers {
		roi, divGrowth, err := getRoiAndDivGrowth(ticker)
		if err != nil {
			if strings.Contains(err.Error(), "Hit the limit") {
				fmt.Println("Retrying...(sleeping for 1 minute)")
				time.Sleep(time.Minute)
				roi, divGrowth, err = getRoiAndDivGrowth(ticker)
			}
			if err != nil {
				if strings.Contains(err.Error(), "Invalid API call") {
					fmt.Println("Invalid ticker", ticker, "consider removing it")
					continue
				} else {
					panic(err)
				}
			}
		}
		// TODO так неправильно считать из-за сплитов!!
		fmt.Println("Dividend growth for " + ticker + " is " + fmt.Sprintf("%.2f", (divGrowth-1)*100))
		printRoi(ticker, divGrowth)
		roiMap[ticker] = roi
	}

	return roiMap
}

func printRoi(ticker string, roi float64) {
	adjRoi := float64(int((roi-1)*10000)) / 100
	adjRoiStr := fmt.Sprintf("%.2f", adjRoi)
	fmt.Println(ticker + " - " + adjRoiStr + "%")
}

func getRoiAndDivGrowth(ticker string) (float64, float64, error) {
	ticker = strings.TrimSpace(ticker)
	var key = os.Getenv("ALPHA_VANTAGE_API_KEY")
	timeForm := "2006-01-02"

	fileName := "files/" + ticker + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Fetching data for", ticker)
			var url = "https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&outputsize=compact&symbol=" + ticker + "&apikey=" + key + "&datatype=json"
			resp, err := http.Get(url)
			if err != nil {
				return 0.0, 0.0, err
			}

			bytes, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return 0.0, 0.0, err
			}

			if strings.Contains(string(bytes), "Thank you for using Alpha Vantage! Please visit https://www.alphavantage.co/premium/ if you would like to have a higher API call volume.") {
				return 0.0, 0.0, errors.New("Hit the limit")
			}
			if strings.Contains(string(bytes), "Invalid API call") {
				return 0.0, 0.0, errors.New("Invalid API call")
			}

			err = ioutil.WriteFile(fileName, bytes, 0644)
			if err != nil {
				return 0.0, 0.0, err
			}

		} else {
			return 0.0, 0.0, err
		}
	}

	var data = &Outer{}
	check(json.Unmarshal(bytes, &data))

	var keys = make([]string, 0)
	stocks := 1.0
	dividends := make([]float64, 0)

	for k, v := range data.Data {
		keys = append(keys, k)
		div, err := strconv.ParseFloat(v.Dividend, 64)
		if err != nil {
			return 0.0, 0.0, err
		}

		if div != 0.0 {
			stockPrice, err := strconv.ParseFloat(v.Close, 64)
			if err != nil {
				return 0.0, 0.0, err
			}
			dividends = append(dividends, div)
			stocks += div / stockPrice * (1 - TAX)
		}
	}

	sort.Strings(keys)

	if len(keys) == 0 {
		return 0.0, 0.0, errors.New("Empty data. Probably hit the limit")
	}

	firstDateStr := keys[0]
	lastDateStr := keys[len(keys)-1]
	firstDate, err := time.Parse(timeForm, firstDateStr)
	if err != nil {
		return 0.0, 0.0, err
	}

	lastDate, err := time.Parse(timeForm, lastDateStr)
	if err != nil {
		return 0.0, 0.0, err
	}

	years := lastDate.Sub(firstDate).Hours() / YEAR_HOURS

	firstValue, err := strconv.ParseFloat(data.Data[firstDateStr].Close, 64)
	lastPrice, err := strconv.ParseFloat(data.Data[lastDateStr].Close, 64)
	lastValue := stocks * lastPrice
	roi := math.Pow(lastValue/firstValue, 1/float64(years))

	divGrowth := 0.0
	if len(dividends) > 1 {
		divGrowth = math.Pow(dividends[0]/dividends[len(dividends)-1], 1/float64(years))
	}

	return roi, divGrowth, nil
}

type Outer struct {
	Meta map[string]string `json: "Meta Data"`
	Data map[string]struct {
		Open     string `json:"1. open"`
		High     string `json:"2. high"`
		Low      string `json:"3. low"`
		Close    string `json:"4. close"`
		AdjClose string `json:"5. adjusted close"`
		Volume   string `json:"6. volume"`
		Dividend string `json:"7. dividend amount"`
	} `json:"Monthly Adjusted Time Series"`
}

func check(err error) {

	if err != nil {
		panic(err)
	}
}
