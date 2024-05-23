package geojson

type (
	Pagination struct {
		PageSize   int `json:"page_size"`
		PageNumber int `json:"page_number"`
		TotalPages int `json:"total_pages"`
	}

	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float32 `json:"coordinates"`
	}

	GeoJsonFeature struct {
		Type       string      `json:"type"`
		Geometry   Geometry    `json:"geometry"`
		Properties interface{} `json:"properties"`
	}

	GeoJsonCollection struct {
		Type       string         `json:"type"`
		Features   GeoJsonFeature `json:"features"`
		Pagination Pagination     `json:"pagination"`
	}
)

func NewFeature(x float32, y float32, properties interface{}) *GeoJsonFeature {
	return &GeoJsonFeature{
		Type: "Feature",
		Geometry: Geometry{
			Type:        "Point",
			Coordinates: []float32{x, y},
		},
		Properties: properties,
	}
}

func NewCollection(features *GeoJsonFeature, pagination Pagination) *GeoJsonCollection {
	return &GeoJsonCollection{
		Type:       "FeatureCollection",
		Features:   *features,
		Pagination: pagination,
	}
}
