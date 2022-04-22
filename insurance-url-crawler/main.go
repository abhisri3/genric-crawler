package main

import (
	"fmt"
	"github.com/gocolly/colly"

)
func main() {
	DataMigration()


	//Base urls
	BaseURL := []string{
		"https://www.aegonlife.com/",
		"https://www.avivaindia.com/",
		"https://www.bhartiaxa.com/",
		"https://www.adityabirlacapital.com/",
		"https://www.canarahsbclife.com/index.html",
		"https://www.exidelife.in/",
	}

	// fmt.Println(BaseURL)

	for _, entry := range BaseURL {
		getUrls(entry)
		
	}
}

func getUrls(entry string){
	c := colly.NewCollector()

	// d := c.Clone()
	cnt :=0
	c.OnHTML("header", func(e *colly.HTMLElement) {

		e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
			cnt++
			tempdata := SearchURLs{}
			tempdata.BaseURL = entry
			link := kf.Attr("href")
			tempdata.Url = kf.Request.AbsoluteURL(link)
			// fmt.Println(link)
			// if err := DB.Where("Url = ?", tempdata.Url).First(&tempdata).Error; err != nil {
				// error handling...
				
			//   }
			
			var count int64


			DB.Model(&SearchURLs{}).Where("Url = ?", tempdata.Url).Count(&count)

			if count == 0 {
				// fmt.Println("))))")
				DB.Create(&tempdata)
			}
				
		})

		
	})


	c.OnHTML("div.header", func(e *colly.HTMLElement) {

		e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
			cnt++
			tempdata := SearchURLs{}
			tempdata.BaseURL = entry
			link := kf.Attr("href")
			tempdata.Url = kf.Request.AbsoluteURL(link)
			// fmt.Println(link)
			// if err := DB.Where("Url = ?", tempdata.Url).First(&tempdata).Error; err != nil {
				// error handling...
				
			//   }
			
			var count int64


			DB.Model(&SearchURLs{}).Where("Url = ?", tempdata.Url).Count(&count)

			if count == 0 {
				// fmt.Println("))))")
				DB.Create(&tempdata)
			}
				
		})

		
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		if cnt == 0 {
			e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
				cnt++
				tempdata := SearchURLs{}
				tempdata.BaseURL = entry
				link := kf.Attr("href")
				tempdata.Url = kf.Request.AbsoluteURL(link)
				// fmt.Println(link)
				// if err := DB.Where("Url = ?", tempdata.Url).First(&tempdata).Error; err != nil {
					// error handling...
					
				//   }
				
				var count int64


				DB.Model(&SearchURLs{}).Where("Url = ?", tempdata.Url).Count(&count)

				if count == 0 {
					// fmt.Println("))))")
					DB.Create(&tempdata)
				}
					
			})
		}
		
	})


	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting : ", r.URL.String())
	})

	c.Visit(entry)

}
