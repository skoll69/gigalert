package main

import (
	"html"
	"strings"
)

var tikettiUrl = "https://www.tiketti.fi/tapahtumat/d/43f50c975c17675bfdf1a1b8019e110f"

type Event struct {
	name     string
	price    string
	date     string
	location string
}

const (
	desc  int = 2
	price int = 4
	date  int = 8
	locId int = 10
	url   int = 15
)

type Location struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type TikettiEvent []string
type Locations map[string]Location

type Tiketti struct {
	Events    []TikettiEvent `json:"events"`
	Locations Locations      `json:"locations"`
}

func getTiketti(searchStrings []string) []Event {
	data := new(Tiketti)
	getJson(tikettiUrl, data, nil)

	out := []Event{}
	for _, searchString := range searchStrings {
		out = append(out, getTikettiOne(searchString, data)...)
	}
	return out
}

func getTikettiOne(text string, data *Tiketti) []Event {
	out := []Event{}
	text = strings.ToLower(text)

	for _, item := range data.Events {
		if strings.Contains(strings.ToLower(item[desc]), text) {
			loc := data.Locations[item[locId]]
			price := strings.Replace(html.UnescapeString(item[price]), "alk. ", "", 1)
			resp := Event{item[desc], price, item[date], loc.Name + ", " + loc.City}
			out = append(out, resp)
		}
	}
	return out
}
