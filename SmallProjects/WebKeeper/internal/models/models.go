package models

import "time"

type Site struct {
	ID   uint `gorm:"primaryKey"`
	Nome string
	URL  string `gorm:"unique"`
	Logs []LogSite
}
type LogSite struct {
	ID         uint `gorm:"primaryKey"`
	SiteID     uint
	Site       Site
	StatusCode int
	LogDes     string
	Hora       time.Time
}
