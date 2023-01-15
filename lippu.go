package main

import (
	"strings"
	"sync"
)

var lippuUrl = "https://public-api.eventim.com/websearch/search/api/exploration/v1/products?language=FI&retail_partner=ADV&search_term="

var headers = [][]string{
	{"origin", "https://www.lippu.fi"},
	{"oidc-client-id", "web__lippu-fi"},
}

type LippuLocation struct {
	Name string `json:"Name"`
	City string `json:"City"`
}

type LiveEntertainment struct {
	StartDate string        `json:"startDate"`
	Location  LippuLocation `json:"location"`
}

type TypeAttributes struct {
	LiveEntertainment LiveEntertainment `json:"liveEntertainment"`
}

type LippuEvent struct {
	Link           string         `json:"link"`
	Name           string         `json:"name"`
	Price          string         `json:"price"`
	TypeAttributes TypeAttributes `json:"typeAttributes"`
}

type Lippu struct {
	Products []LippuEvent `json:"products"`
}

func getLippu(searchStrings []string) []Event {
	var wg sync.WaitGroup

	events := []Event{}
	for _, searchString := range searchStrings {
		wg.Add(1)
		go getLippuOne(searchString, &events, &wg)
	}
	wg.Wait()
	return events
}

func getLippuOne(searchString string, events *[]Event, wg *sync.WaitGroup) {
	println(searchString)
	data := new(Lippu)
	getJson(lippuUrl+searchString, data, headers)

	for _, item := range data.Products {
		attrs := item.TypeAttributes.LiveEntertainment
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchString)) { // Avoid results from fuzzy search
			event := Event{item.Name, item.Price, attrs.StartDate, attrs.Location.Name + ", " + attrs.Location.City}
			*events = append(*events, event)
		}
	}
	wg.Done()
}

// func getLippuOne(searchString string) []Event {
// 	events := []Event{}

// 	data := new(Lippu)
// 	getJson(lippuUrl+searchString, data, headers)

// 	for _, item := range data.Products {
// 		attrs := item.TypeAttributes.LiveEntertainment
// 		event := Event{item.Name, item.Price, attrs.StartDate, attrs.Location.Name + ", " + attrs.Location.City}
// 		// event := Event{item.Name, item.Price, item[date], loc.Name + ", " + loc.City}
// 		events = append(events, event)
// 	}
// 	return events
// }
