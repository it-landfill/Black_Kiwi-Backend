package black_kiwi_mobile_queries

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates, st_distance(poi.geom,st_geogfromtext('POINT(11.3428 44.4939)')) as meters
FROM "black-kiwi_data".poi_list as poi
JOIN "black-kiwi_data".categories as cat on poi.category = cat.id
WHERE cat.name = 'park' and poi.rank<0
ORDER BY meters;
*/
func GetRecommendation(minRank float64, lat float64, lon float64, category string, limit int) (success bool, poiList []black_kiwi_data_structs.PoiItem) {
	success = true
	queryStr := ""
	poiList = []black_kiwi_data_structs.PoiItem{}

	queryStr += fmt.Sprintf("SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates, st_distance(poi.geom,st_geogfromtext('POINT(%f %f)')) as meters\n", lat, lon)
	queryStr += "FROM \"black-kiwi_data\".poi_list as poi\n"
	queryStr += "JOIN \"black-kiwi_data\".categories as cat on poi.category = cat.id\n"

	if category != "" && minRank != 0 {
		queryStr += fmt.Sprintf("WHERE cat.name = '%s' and poi.rank>=%d\n", category, minRank)
	} else {
		if category != "" {
			queryStr += fmt.Sprintf("WHERE cat.name = '%s'\n", category)
		}
		if minRank != 0 {
			queryStr += fmt.Sprintf("WHERE poi.rank>=%d\n", minRank)
		}
	}

	queryStr += "ORDER BY meters\n"

	if limit != 0 {
		queryStr += fmt.Sprintf("LIMIT %d\n", limit)
	}

	queryStr += ";"

	rows, err := black_kiwi_db_utils.ConnPool.Query(context.Background(), queryStr)
	if err != nil {
		log.WithFields(log.Fields{"error":err}).Error("QueryRow failed for recommended POIS.")
		success = false
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var rank float32
		var category string
		var coordinates string
		var meters float64
		err = rows.Scan(&id, &name, &rank, &category, &coordinates, &meters)
		if err != nil {
			log.WithFields(log.Fields{"error":err}).Error("QueryRow failed while scanning rows for recommended POIS.")
			success = false
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

		poiList = append(poiList, poiItem)
	}

	return
}
