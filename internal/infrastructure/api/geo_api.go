package api

import (
	"errors"
	"fmt"
)

const (
	GEO_API_URL = "https://json.geoapi.pt/"
)

var (
	ErrLocalNotFound   = errors.New("there ins't any local with those coordinates")
	ErrZipCodeNotFound = errors.New("there ins't any local with that zip code")
)

var geoApi *Api = New(GEO_API_URL)

type (
	Location struct {
		Lon       float64 `json:"lon"`
		Lat       float64 `json:"lat"`
		Distrito  string  `json:"distrito"`
		Concelho  string  `json:"concelho"`
		Freguesia string  `json:"freguesia"`
		AltitudeM int     `json:"altitude_m"`
		Uso       string  `json:"uso"`
		SEC       string  `json:"SEC"`
		SS        string  `json:"SS"`
		Rua       string  `json:"rua"`
		CP        string  `json:"CP"`
	}

	Street struct {
		Arteria string `json:"Artéria"`
		Local   string `json:"Local"`
		Troco   string `json:"Troço"`
		Porta   string `json:"Porta"`
		Cliente string `json:"Cliente"`
	}

	Point struct {
		ID          string    `json:"id"`
		Rua         string    `json:"rua"`
		Casa        string    `json:"casa"`
		Coordenadas []float64 `json:"coordenadas"`
	}

	CPHousing struct {
		CP               string      `json:"CP"`
		CP4              string      `json:"CP4"`
		CP3              string      `json:"CP3"`
		Distrito         string      `json:"Distrito"`
		Concelho         string      `json:"Concelho"`
		Localidade       string      `json:"Localidade"`
		DesignacaoPostal string      `json:"Designação Postal"`
		Partes           []Street    `json:"partes"`
		Pontos           []Point     `json:"pontos"`
		Ruas             []string    `json:"ruas"`
		Centro           []float64   `json:"centro"`
		Poligono         [][]float64 `json:"poligono"`
		Centroide        []float64   `json:"centroide"`
		CentroDeMassa    []float64   `json:"centroDeMassa"`
		Municipio        string      `json:"municipio"`
		CodigoINE        struct {
			CodigoINEDistrito  string `json:"codigoineDistrito"`
			CodigoINEMunicipio string `json:"codigoineMunicipio"`
		} `json:"codigoine"`
	}
)

func GetLocation(x, y float32) (*Location, error) {
	var location Location

	if err := geoApi.getJson(fmt.Sprintf("gps/%f,%f", x, y), &location); err != nil {
		return nil, ErrLocalNotFound
	}

	return &location, nil
}

func GetCPHousing(zipCode string) (*CPHousing, error) {
	var cp CPHousing

	if err := geoApi.getJson(fmt.Sprintf("cp/%s", zipCode), &cp); err != nil {
		return nil, ErrZipCodeNotFound
	}

	return &cp, nil
}
