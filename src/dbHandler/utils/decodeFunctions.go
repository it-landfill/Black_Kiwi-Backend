package black_kiwi_db_utils

import (
	"context"
	"encoding/json"

	"ITLandfill/Black-Kiwi/structs/data_structs"

	log "github.com/sirupsen/logrus"
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

func GetIDFromCategory() *map[black_kiwi_data_structs.Categories]int {

	catMap := map[black_kiwi_data_structs.Categories]int{}

	rows, err := ConnPool.Query(context.Background(), "SELECT id, name FROM \"black-kiwi_data\".categories;")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Query failed while getting categories.")
		return nil
	}

	for rows.Next() {
		var categoryName string
		var id int
		err = rows.Scan(&id, &categoryName)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Scan failed while getting categories.")
			return nil
		}

		catMap[StringToCategory(categoryName)] = id
	}

	return &catMap
}