package black_kiwi_db_utils

import "encoding/json"

func JSONtoCoordinates(jsonStr string) (lat float64, lon float64) {
	type jsonCoordinates struct {
		Type string `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	var coordinates jsonCoordinates
	json.Unmarshal([]byte(jsonStr), &coordinates)
	coords := coordinates.Coordinates
	return coords[0], coords[1]
}