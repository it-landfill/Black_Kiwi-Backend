package black_kiwi_admin_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT quartieri.nomequart, quartieri.geom, count(poi.name) as poi_density
FROM "black-kiwi_data".poi_list as poi
JOIN "black-kiwi_data"."quartieri-bologna" as quartieri on st_within(poi.geom, quartieri.geom)
GROUP BY quartieri.nomequart, quartieri.geom;

SELECT quartieri.nomequart, quartieri.geom, count(poi.name) as poi_density FROM "black-kiwi_data".poi_list as poi JOIN "black-kiwi_data"."quartieri-bologna" as quartieri on st_within(poi.geom, quartieri.geom) GROUP BY quartieri.nomequart, quartieri.geom;
*/
func GetPOIQuartieri() (result bool, poiList *[]black_kiwi_data_structs.QuartiereInfo)  {

	rows, err := black_kiwi_db_utils.ConnPool.Query(context.Background(), "SELECT quartieri.nomequart, quartieri.geom, count(poi.name) as poi_density FROM \"black-kiwi_data\".poi_list as poi JOIN \"black-kiwi_data\".\"quartieri-bologna\" as quartieri on st_within(poi.geom, quartieri.geom) GROUP BY quartieri.nomequart, quartieri.geom;")
	if err != nil {
		log.WithFields(log.Fields{"error":err}).Error("QueryRow failed for get POI quartieri.")
		return false, nil
	}
	defer rows.Close()

	quartieriList := []black_kiwi_data_structs.QuartiereInfo{}

	for rows.Next() {
		var name string
		var geometry string
		var density int
		err = rows.Scan(&name, &geometry, &density)
		if err != nil {
			if (err.Error() == "no rows in result set") {
				log.Info("No Quartieri found in db.")
				return true, nil
			}

			log.WithFields(log.Fields{"error":err}).Error("QueryRow failed while scanning rows for POI quartieri.")
			return false, nil
		}

		quartiereInfo := black_kiwi_data_structs.QuartiereInfo{
			Name: name,
			Geometry: geometry,
			Density: density,
		}

		quartieriList = append(quartieriList, quartiereInfo)
	}
	return true, &quartieriList
}