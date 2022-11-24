package black_kiwi_admin_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT json_build_object('type', 'FeatureCollection', 'features', json_agg(ST_AsGeoJSON(quartieri.*)::json))
FROM (SELECT quartieri.name as name, quartieri.geom, count(req.id) as value
      FROM "black-kiwi_data"."Quartieri" as quartieri
               LEFT JOIN "black-kiwi_data"."Requests" as req on st_within(req.geom, quartieri.geom)
      GROUP BY quartieri.name, quartieri.geom) as quartieri;

SELECT json_build_object('type', 'FeatureCollection', 'features', json_agg(ST_AsGeoJSON(quartieri.*)::json)) FROM (SELECT quartieri.name as name, quartieri.geom, count(req.id) as value FROM "black-kiwi_data"."Quartieri" as quartieri LEFT JOIN "black-kiwi_data"."Requests" as req on st_within(req.geom, quartieri.geom) GROUP BY quartieri.name, quartieri.geom) as quartieri;
 */
func GetCheckinQuartieri() (result bool, featureCollection string)  {

	err := black_kiwi_db_utils.ConnPool.QueryRow(context.Background(), "SELECT json_build_object('type', 'FeatureCollection', 'features', json_agg(ST_AsGeoJSON(quartieri.*)::json)) FROM (SELECT quartieri.name as name, quartieri.geom, count(req.id) as value FROM \"black-kiwi_data\".\"Quartieri\" as quartieri LEFT JOIN \"black-kiwi_data\".\"Requests\" as req on st_within(req.geom, quartieri.geom) GROUP BY quartieri.name, quartieri.geom) as quartieri;").Scan(&featureCollection)
	if err != nil {
		log.WithFields(log.Fields{"error":err}).Error("QueryRow failed for get POI quartieri.")
		return false, ""
	}

	return true, featureCollection
}