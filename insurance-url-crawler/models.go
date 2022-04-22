package main

import "gorm.io/gorm"

type SearchURLs struct {
	gorm.Model
	Url         string
	BaseURL 	string
	Type 		string //wheteher plan type other type or null
}