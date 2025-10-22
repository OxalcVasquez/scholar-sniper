package main

import (
	"fmt"

	"github.com/oxalcvasquez/scholar-sni/db"
	"github.com/oxalcvasquez/scholar-sni/scraper"
)

func main() {

	db.InitDB()

	url := "https://www.kth.se/en/studies/master/computer-science?utm_source=chatgpt.com"

	data, err := scraper.ScrapeScholarships(url)

	if err != nil {
		fmt.Println("Error scraping:", err)
		return
	}

	fmt.Println("Title:", data.Title)
	fmt.Println("Deadline:", data.Deadline)
	fmt.Println("Requirements:", data.Requirements)

}
