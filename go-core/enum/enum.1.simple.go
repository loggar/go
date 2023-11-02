package brand

// Brand is the type of the app
type Brand uint8

const (
	// NIKE represents the NIKE brand.
	NIKE Brand = iota
	// ADIDAS represents the ADIDAS brand.
	ADIDAS
	// PUMA represents the PUMA brand.
	PUMA
)

var brandStrings = map[Brand]string{
	NIKE:   "NIKE",
	ADIDAS: "ADIDAS",
	PUMA:   "PUMA",
}

// String returns the string representation of the brand.
func (b Brand) String() string {
	if brand, ok := brandStrings[b]; ok {
		return brand
	}
	return "UNKNOWN"
}

// AllBrands returns a slice of all available brands.
func AllBrands() []Brand {
	return []Brand{NIKE, ADIDAS, PUMA}
}

// IsBrand checks if a given string matches any known brand.
func IsBrand(s string) bool {
	for _, brand := range AllBrands() {
		if brand.String() == s {
			return true
		}
	}
	return false
}
