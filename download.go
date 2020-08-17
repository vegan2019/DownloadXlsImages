package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	isbns := ArrayFromCSV("./excel/Complete-HUP.csv")
	//fmt.Println(isbns)

	lenISBN := len(isbns)
	fmt.Println(lenISBN)
	//去掉第一个元素，因为第一个不是
	//isbns = isbns[1 : lenISBN-1]
	isbns = isbns[1:3]
	for _, isbn := range isbns {
		//https://www.hup.harvard.edu/images/jackets/9780674066649-lg.jpg
		fileUrl := "https://www.hup.harvard.edu/images/jackets/" + isbn + "-lg.jpg"

		err := DownloadFile("./HUB-images/"+isbn+"-lg.jpg", fileUrl)

		if err != nil {
			panic(err)
		}

		fmt.Println("Downloaded: " + fileUrl)
		time.Sleep(2000 * time.Millisecond)
	}

}

func oldmain() {
	fileUrl := "https://golangcode.com/logo.svg"
	err := DownloadFile("./HUB-images/logo.svg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func ArrayFromCSV(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var item string
	var items []string

	for _, record := range records {

		item = record[1]
		items = append(items, item)
	}

	return items
}
