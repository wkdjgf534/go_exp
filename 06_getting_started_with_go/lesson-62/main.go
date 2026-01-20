package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const URL = "https://bored-api.appbrewery.com/random"

type boringResponse struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility string  `json:"accessibility"`
	Duration      string  `json:"duration"`
	KindFriendly  bool    `json:"kidFriendly"`
}

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}

	resp, err := client.Do(req)
	fmt.Println(resp.Body)
	if err != nil {
		fmt.Println("error doing request: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error invalid status code:", resp.Status)
		return
	}

	var boringRes boringResponse
	if err := json.NewDecoder(resp.Body).Decode(&boringRes); err != nil {
		fmt.Println("error reading json response body: ", err)
		return
	}

	fmt.Printf("%+v\n", boringRes)
	// {Activity:Clean out your car Type:busywork Participants:1 Price:0 Link: Key:2896176 Accessibility:Minor challenges Duration:minutes KindFriendly:true}
}
