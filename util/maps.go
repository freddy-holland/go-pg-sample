package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"fholl.net/go-pg-sample/models"
)

var key string = "***"

type DistanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Distance struct {
				Text  string  `json:"text"`
				Value float32 `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string  `json:"text"`
				Value float32 `json:"value"`
			} `json:"duration"`
		} `json:"elements"`
	} `json:"rows"`
}

type ContractorDMR struct {
	Contractor *models.Contractor
	Elements   *DistanceMatrixResponse
}

func GetDistance(destination string, origin string) (*DistanceMatrixResponse, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?destinations=%s&origins=%s&units=metric&key=%s", url.QueryEscape(destination), url.QueryEscape(origin), key)
	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var dmr DistanceMatrixResponse
	err = json.Unmarshal(body, &dmr)
	if err != nil {
		return nil, err
	}

	if len(dmr.Rows) == 0 || len(dmr.Rows[0].Elements) == 0 {
		return nil, fmt.Errorf("No data found")
	}

	return &dmr, nil
}
