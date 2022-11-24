package black_kiwi_admin_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
INSERT INTO "black-kiwi_data"."Pois" (id, name, category, rank, geom)
VALUES (DEFAULT, 'AjejeBa', 4, -1, ST_SetSRID(ST_MakePoint(-71.1043443253471, 42.3150676015829),4326))
RETURNING "id";

INSERT INTO "black-kiwi_data"."Pois" (id, name, category, rank, geom) VALUES (DEFAULT, $3, $4, $5, ST_SetSRID(ST_MakePoint($1, $2),4326)) RETURNING "id";
*/
func AddPOI(poi black_kiwi_data_structs.NewPoiItem) int  {
	log.WithField("New POI", poi).Info("AddPOI query called")

	catMap := black_kiwi_db_utils.GetIDFromCategory()

	catID := (*catMap)[poi.Category]
	var poiID int
	err := black_kiwi_db_utils.ConnPool.QueryRow(context.Background(), "INSERT INTO \"black-kiwi_data\".\"Pois\" (id, name, category, rank, geom) VALUES (DEFAULT, $3, $4, $5, ST_SetSRID(ST_MakePoint($1, $2),4326)) RETURNING \"id\";", poi.Coord.Longitude, poi.Coord.Latitude, poi.Name, catID, poi.Rank).Scan(&poiID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("QueryRow failed while adding POI.")
		return -1
	}

	log.WithFields(log.Fields{"POI ID": poiID}).Debug("QueryRow succeeded while adding POI.")

	return poiID
}