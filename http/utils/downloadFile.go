package utils

import "time"
import "fmt"
import "os"
import "github.com/cavaliercoder/grab"

func Download(dest string, url string) string{

if stat, e := os.Stat(dest); os.IsNotExist(e){
	fmt.Printf("File/directory %s does not exist. Creating.", dest)
	fmt.Println()

	os.Mkdir(wd + "/" + dest, 0777)

	stat2, _ := os.Stat(dest)
	fmt.Printf("Created %s. Is directory - %s", dest, stat2.IsDir())
	fmt.Println()

} else {
	fmt.Printf("File/directory %s already exists. Is directory - %s", dest, stat.IsDir())
	fmt.Println()
}


resp, _ := grab.Get(dest, url)
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
				fmt.Println("response - ", resp.HTTPResponse)
				os.Exit(1)
			}
			fmt.Printf("Successfully downloaded to %s\n", resp.Filename)
			return resp.Filename
		}
	}
}