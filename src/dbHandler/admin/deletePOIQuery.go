package black_kiwi_admin_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
DELETE FROM "black-kiwi_data"."Pois" WHERE id = 32*/
func DeletePOI(poiID int) bool {
	log.WithField("POI ID", poiID).Info("DeletePOI query called")

	_, err := black_kiwi_db_utils.ConnPool.Exec(context.Background(), "DELETE FROM \"black-kiwi_data\".\"Pois\" WHERE id = $1;", poiID) //TODO: Cascade delete on requests?
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Exec failed while deleting POI.")
		return false
	}

	log.WithFields(log.Fields{"POI ID": poiID}).Debug("Exec succeeded while deleting POI.")

	return true
}