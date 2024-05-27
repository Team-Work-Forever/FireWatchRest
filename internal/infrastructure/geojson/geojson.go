package geojson

import "github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/pagination"

type (
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	GeoJsonFeature struct {
		Type       string      `json:"type"`
		Geometry   Geometry    `json:"geometry"`
		Properties interface{} `json:"properties"`
	}

	GeoJsonCollection struct {
		Type       string                `json:"type"`
		Features   []GeoJsonFeature      `json:"features"`
		Pagination pagination.Pagination `json:"pagination"`
	}
)

func NewFeature(x float64, y float64, properties interface{}) *GeoJsonFeature {
	return &GeoJsonFeature{
		Type: "Feature",
		Geometry: Geometry{
			Type:        "Point",
			Coordinates: []float64{x, y},
		},
		Properties: properties,
	}
}

func NewCollection(features []GeoJsonFeature, pagination pagination.Pagination) *GeoJsonCollection {
	return &GeoJsonCollection{
		Type:       "FeatureCollection",
		Features:   features,
		Pagination: pagination,
	}
}
