package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	Name          string    `json:"name"`
	Pi            float64   `json:"pi"`
	TheBestNumber int       `json:"best_number"`
	RightNow      time.Time `json:"right_now"`
}

func payload() (*Payload, error) {
	req, _ := http.NewRequest("GET", "http://localhost:9292", nil)
	req.Header.Set("Accepts", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	// Parse the response
	var payload Payload
	if err := json.Unmarshal(content, &payload); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &payload, nil
}

func main() {
	payload, _ := payload()
	fmt.Printf("%+v\n", payload)
}
