package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

//Fake DB

var File *os.File

func main() {

	DataMigration()

	// Searchdb := []Search{
	// 	{"https://www.aegonlife.com/insurance-plans/retirement-plans/aegon-life-insta-pension-plan", "h1.node__title > span", "div.star-insurance-plan > div.desc-block > p", "div.carousel-content > div > p", "div.carousel-content > div > p"},
	// 	{"https://www.aegonlife.com/insurance-plans/saving-plans/aegon-life-jeevan-riddhi-insurance-plan", "h1.node__title > span", "div.product-buy-plan > div > div.desc-block > p", "div.carousel-content > p", "div.carousel-content > p"},
	// 	{"https://www.bhartiaxa.com/savings-plans/guaranteed-wealth-pro", "h1.banner-main-title", "p.titel_desc", "div.rowImageWrapper > p", "tr > td"},
	// 	{"https://www.bhartiaxa.com/term-insurance/flexi-term-pro", "h1.banner-main-title", "p.titel_desc", "div.rowImageWrapper > p.rowImageTitle", "tr td"},
	// }

	var Searchdb []Search
	DB.Find(&Searchdb)
	File, _ = os.Create("Insurances.txt")
	for _, entry := range Searchdb {
		getInsuranceDetails(entry)
		time.Sleep(2 * time.Second)
	}

}

func getInsuranceDetails(entry Search) {
	c := colly.NewCollector()

	// d := c.Clone()
	tempdata := InsuranceData{}
	tempdata.Benefits = ""
	tempdata.Eligibility = ""
	c.OnHTML("body", func(e *colly.HTMLElement) {

		tempdata.Title = e.ChildText(entry.Title)
		tempdata.Desc = e.ChildText(entry.Desc)
		e.ForEach(entry.Benefits, func(_ int, elem *colly.HTMLElement) {
			tempdata.Benefits = tempdata.Benefits + elem.Text + "\n"
		})
		e.ForEach(entry.Eligibility, func(_ int, elem *colly.HTMLElement) {
			tempdata.Eligibility = tempdata.Eligibility + elem.Text + "\n"
		})
		tempdata.Price = e.ChildText(entry.Price)

		// js, err := json.MarshalIndent(tempdata, "", "    ")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		DB.Create(&tempdata)

		// File.WriteString(string(js) + "\n")

		// fmt.Println(string(js))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting : ", r.URL.String())
	})

	c.Visit(entry.Url)

}
