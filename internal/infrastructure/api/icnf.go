package api

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

const (
	ICFN_URL = "https://fogos.icnf.pt:8443/localizador/"
)

var icfnApi *Api = New(ICFN_URL)

type (
	ICNF struct {
		XMLName       xml.Name `xml:"Response"`
		Distrito      string   `xml:"DISTRITO"`
		Concelho      string   `xml:"CONCELHO"`
		Freguesia     string   `xml:"FREGUESIA"`
		Local         string   `xml:"LOCAL"`
		INE           string   `xml:"INE"`
		NUT           string   `xml:"NUT"`
		X             int      `xml:"X"`
		Y             int      `xml:"Y"`
		Lat           string   `xml:"Lat"`
		Lon           string   `xml:"Lon"`
		TipoOperacao  string   `xml:"TipoOperacao"`
		DataOperacao  string   `xml:"DataOperacao"`
		RCM           int      `xml:"RCM"`
		Percentil15d  int      `xml:"Percentil_15d"`
		PercentilAno  int      `xml:"Percentil_ano"`
		Perigosidade  float64  `xml:"Perigosidade"`
		FWI           int      `xml:"FWI"`
		Legalidade    string   `xml:"Legalidade"`
		RiscoOperacao int      `xml:"RiscoOperacao"`
		Resposta      string   `xml:"Resposta"`
		Analise       int      `xml:"Analise"`
	}
)

func GetICNFIndex(x, y float64, has_aid_team bool) (*ICNF, error) {
	var index ICNF
	var aidString string

	if has_aid_team {
		aidString = "Sim"
	} else {
		aidString = "Não"
	}

	aid := url.QueryEscape(aidString)
	url := fmt.Sprintf("webservicequeimadas2019.asp?Operacao=Queima&data=2024-05-25&lat=%.16g&lon=%.16g&apoio=%s&tecnicoFC=DEFAULT", x, y, aid)

	if err := icfnApi.getXml(url, &index); err != nil {
		return nil, err
	}

	return &index, nil
}
