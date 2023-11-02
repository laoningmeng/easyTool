package model

import "time"

type BpmSettings struct {
	Uuid      string
	Name      string
	FormKey   string
	AppId     string
	DataCode  string
	Fields    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
