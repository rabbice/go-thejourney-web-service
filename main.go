package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)

	APIKey = "[GOOGLE_API_KEY]"

	mux := http.NewServeMux()

	mux.HandleFunc("/journeys", cors(getJourneys))

	mux.HandleFunc("/recommendations", cors(getRecommendations))

	fileServer := http.FileServer(http.Dir("public/static"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting the API on :...")
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
