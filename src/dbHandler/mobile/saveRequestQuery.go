package black_kiwi_mobile_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
)

/*
INSERT INTO "black-kiwi_data"."Requests" (id, timestamp, category, rank, geom) VALUES (DEFAULT, DEFAULT, $1, $2, ST_SetSRID(ST_MakePoint($3, $4),4326)))
*/
func SaveRequest(minRank float64, lat float64, lon float64, category string) bool {
	log.WithFields(log.Fields{"minRank": minRank, "lat": lat, "lon": lon, "category": category}).Info("SaveRequest query called")

	catID := (*black_kiwi_db_utils.CatMap)[black_kiwi_db_utils.StringToCategory(category)]

	_, err := black_kiwi_db_utils.ConnPool.Exec(context.Background(), "INSERT INTO \"black-kiwi_data\".\"Requests\" (id, timestamp, category, rank, geom) VALUES (DEFAULT, DEFAULT, $1, $2, ST_SetSRID(ST_MakePoint($3, $4),4326)));", catID, minRank, lon, lat)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Exec failed while saving request.")
		return false
	}

	log.WithFields(log.Fields{"minRank": minRank, "lat": lat, "lon": lon, "category": category}).Debug("Exec succeeded at saving request.")

	return true
}