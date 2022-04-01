package main

// Query is the user representation of a query to build recommendations.
type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

type Journey struct {
	Name   string
	Places []string
}

// Place is the Google Place model representation.
type Place struct {
	Geometry *googleGeometry `json:"geometry"`
	Name     string          `json:"name"`
	Icon     string          `json:"icon"`
	Photos   []*googlePhoto  `json:"photos"`
	Vicinity string          `json:"vicinity"`
}

type googleResponse struct {
	Results []Place `json:"results"`
}

type googleGeometry struct {
	Location googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

// CostRange represents a range between Cost values.
type CostRange struct {
	From Cost
	To   Cost
}
