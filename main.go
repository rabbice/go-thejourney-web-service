package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)

	APIKey = "AIzaSyAZcw_SdmOmqG0LTe3fPkbAMBA8dXFWcGI"

	mux := http.NewServeMux()

	mux.HandleFunc("/journeys", cors(getJourneys))

	mux.HandleFunc("/recommendations", cors(getRecommendations))

	fileServer := http.FileServer(http.Dir("public/static"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	log.Println("Starting the API on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
