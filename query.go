package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// APIKey Google API Key.
var APIKey string

func (q *Query) find(types string) (gR *googleResponse, err error) {
	urlString := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

	vals := make(url.Values)

	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("type", types)
	vals.Set("key", APIKey)

	if q.CostRangeStr == "" {
		r := ParseCostRange(q.CostRangeStr)

		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}

	res, err := http.Get(urlString + "?" + vals.Encode())
	if err != nil {
		return
	}
	defer res.Body.Close()

	var response googleResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	gR = &response
	return
}

// Run runs the query concurrently, and returns the results.
func (q *Query) Run() (results []interface{}) {
	rand.Seed(time.Now().UnixNano())

	var w sync.WaitGroup
	var l sync.Mutex

	places := make([]interface{}, len(q.Journey))

	for i, typeJourney := range q.Journey {
		w.Add(1)

		go func(types string, i int) {
			defer w.Done()

			response, err := q.find(types)
			if err != nil {
				log.Println("Failed to find places:", err)
				return
			}

			if len(response.Results) == 0 {
				log.Println("No places found for", types)
				return
			}

			for _, result := range response.Results {
				for _, photo := range result.Photos {
					photoURL := fmt.Sprintf(
						"https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photoreference=%s&key=%s",
						photo.PhotoRef,
						APIKey,
					)

					photo.URL = photoURL
				}
			}

			randI := rand.Intn(len(response.Results))

			l.Lock()
			places[i] = response.Results[randI]
			l.Unlock()
		}(typeJourney, i)
	}

	w.Wait()

	results = places
	return
}
