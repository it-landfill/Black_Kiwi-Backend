package black_kiwi_default_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates
FROM "black-kiwi_data".poi_list as poi
JOIN "black-kiwi_data".categories as cat on poi.category = cat.id;
*/
func GetPOIS() (result bool, poiList *[]black_kiwi_data_structs.PoiItem)  {

	rows, err := black_kiwi_db_utils.ConnPool.Query(context.Background(), "SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates FROM \"black-kiwi_data\".poi_list as poi JOIN \"black-kiwi_data\".categories as cat on poi.category = cat.id")
	if err != nil {
		log.WithFields(log.Fields{"error":err}).Error("QueryRow failed for get POIS.")
		return false, nil
	}
	defer rows.Close()

	tmpList := []black_kiwi_data_structs.PoiItem{}

	for rows.Next() {
		var id int
		var name string
		var rank float32
		var category string
		var coordinates string
		err = rows.Scan(&id, &name, &rank, &category, &coordinates)
		if err != nil {
			if (err.Error() == "no rows in result set") {
				log.Info("No POIs found in db.")
				return true, nil
			}

			log.WithFields(log.Fields{"error":err}).Error("QueryRow failed while scanning rows for get POIS.")
			return false, nil
		}

		tmpLon, tmpLat := black_kiwi_db_utils.JSONtoCoordinates(coordinates)

		poiItem := black_kiwi_data_structs.PoiItem{
			Id:       id,
			Name:     name,
			Rank:     rank,
			Category: black_kiwi_db_utils.StringToCategory(category),
			Coord: black_kiwi_data_structs.Coordinates{
				Latitude:  tmpLat,
				Longitude: tmpLon,
			},
		}

		tmpList = append(tmpList, poiItem)
	}
	return true, &tmpList
}