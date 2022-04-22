package main

import "gorm.io/gorm"



type Search struct {
	gorm.Model
	Url         string
	Title       string
	Desc        string
	Benefits    string
	Eligibility string
	Price 		string
}

type InsuranceData struct {
	gorm.Model
	Title       string
	Desc        string
	Benefits    string
	Eligibility string
	Price 		string
}
