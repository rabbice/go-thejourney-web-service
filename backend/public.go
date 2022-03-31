package thejourney

// Facade matches with types with public versions.
type Facade interface {
	Public() interface{}
}

// Public return the public version of a type,
// or the same type if it don't have a public version.
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}

	return o
}

// Public returns the public version of Place.
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Geometry.Location.Lat,
		"lng":      p.Geometry.Location.Lng,
	}
}
