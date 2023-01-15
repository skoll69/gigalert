package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func getJson(url string, target interface{}, headers [][]string) error {
	var client = &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	for _, header := range headers {
		req.Header.Set(header[0], header[1])
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// bodyBytes, _ := io.ReadAll(res.Body)
	// fmt.Println(string(bodyBytes))
	return json.NewDecoder(res.Body).Decode(target)
}
