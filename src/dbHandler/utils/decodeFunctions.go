package black_kiwi_db_utils

import (
	"ITLandfill/Black-Kiwi/structs/data_structs"
	"encoding/json"
)

func JSONtoCoordinates(jsonStr string) (float64, float64) {
	type jsonCoordinates struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	var coordinates jsonCoordinates
	json.Unmarshal([]byte(jsonStr), &coordinates)
	coords := coordinates.Coordinates
	return coords[0], coords[1]
}

func StringToCategory(catStr string) black_kiwi_data_structs.Categories {
	switch catStr {
	case "Park":
		return black_kiwi_data_structs.PARK
	case "Museum":
		return black_kiwi_data_structs.MUSEUM
	case "Historical Building":
		return black_kiwi_data_structs.HISTORICAL_BUILDING
	case "Theater":
		return black_kiwi_data_structs.THEATER
	case "Department":
		return black_kiwi_data_structs.DEPARTMENT
	}
	return black_kiwi_data_structs.PARK
}
