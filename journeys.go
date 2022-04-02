package main

import "strings"

func (j *Journey) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.Places, "|"),
	}
}

// Journeys are static recommendations
var Journeys = []interface{}{
	&Journey{
		Name: "Romantic",
		Places: []string{
			"park", "bar", "movie_theater", "restaurant", "florist",
		},
	},
	&Journey{
		Name: "Shopping",
		Places: []string{
			"department_store", "convenience_store", "cafe", "clothing_store", "jewelry_store", "shoe_store", "supermarket",
		},
	},
	&Journey{
		Name: "Night Out",
		Places: []string{
			"bar", "casino", "food", "night_club", "bar",
		},
	},
	&Journey{
		Name: "Culture",
		Places: []string{
			"museum", "cafe", "cemetery", "library", "art_gallery",
		},
	},
	&Journey{
		Name: "Pamper",
		Places: []string{
			"hair_care", "beauty_salon", "cafe", "spa",
		},
	},
}
