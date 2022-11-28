package black_kiwi_admin_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
UPDATE "black-kiwi_data"."Pois" SET name = 'fwe', category = 4, rank = 3 WHERE id = 29
*/
func EditPOI(poi black_kiwi_data_structs.PoiItem) bool  {
	log.WithField("Edited POI", poi).Info("EditPOI query called")

	catID := (*black_kiwi_db_utils.CatMap)[poi.Category]

	_, err := black_kiwi_db_utils.ConnPool.Exec(context.Background(), "UPDATE \"black-kiwi_data\".\"Pois\" SET name = $1, category = $2, rank = $3 WHERE id = $4;", poi.Name, catID, poi.Rank, poi.Id)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Exec failed while editing POI.")
		return false
	}

	log.WithFields(log.Fields{"Edited POI": poi}).Debug("Exec succeeded while editing POI.")

	return true
}