package black_kiwi_admin_queries

import (
	"context"
	"time"
	"fmt"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT req.timestamp, cat.name as category, req.rank, st_asgeojson(req.geom) as coordinates, st_clusterkmeans(req.geom, 2) OVER() as cid
FROM "black-kiwi_data"."Requests" as req
JOIN "black-kiwi_data"."Categories" as cat on req.category = cat.id
WHERE req.timestamp > timestamp '2022-10-01' and req.timestamp < timestamp '2022-11-01';
*/
func GetClusterLocation(startDate *time.Time, endDate *time.Time, nCluster int, radius float32) (result bool, reqList *[]black_kiwi_data_structs.ClusterInfo)  {
	log.WithFields(log.Fields{"startDate": startDate, "endDate": endDate, "nCluster": nCluster, "radius": radius}).Info("GetClusterLocationQuery called")

	var queryStr string
	if radius == 0 {
		queryStr = fmt.Sprintf("SELECT req.timestamp, cat.name, req.rank, st_asgeojson(req.geom), st_clusterkmeans(req.geom, %d) OVER() as cid FROM \"black-kiwi_data\".\"Requests\" as req JOIN \"black-kiwi_data\".\"Categories\" as cat on req.category = cat.id", nCluster)
	} else {
		queryStr = fmt.Sprintf("SELECT req.timestamp, cat.name, req.rank, st_asgeojson(req.geom), st_clusterkmeans(req.geom, %d, %f) OVER() as cid FROM \"black-kiwi_data\".\"Requests\" as req JOIN \"black-kiwi_data\".\"Categories\" as cat on req.category = cat.id", nCluster, radius)
	}

	if startDate != nil && endDate != nil {
		queryStr += fmt.Sprintf(" WHERE req.timestamp > timestamp '%s' and req.timestamp < timestamp '%s'", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	} 
	
	queryStr += ";"
	
	rows, err := black_kiwi_db_utils.ConnPool.Query(context.Background(), queryStr)
	if err != nil {
		log.WithFields(log.Fields{"error":err}).Error("QueryRow failed for get Request Locations.")
		return false, nil
	}
	defer rows.Close()

	tmpList := []black_kiwi_data_structs.ClusterInfo{}

	for rows.Next() {
		var timestamp time.Time
		var rank float64
		var category string
		var coordinates string
		var cid int
		err = rows.Scan(&timestamp, &category, &rank, &coordinates, &cid)
		if err != nil {
			if (err.Error() == "no rows in result set") {
				log.Info("No Request found in db.")
				return true, nil
			}

			log.WithFields(log.Fields{"error":err}).Error("QueryRow failed while scanning rows for get Clustering Locations.")
			return false, nil
		}

		tmpLon, tmpLat := black_kiwi_db_utils.JSONtoCoordinates(coordinates)

		reqItem := black_kiwi_data_structs.ClusterInfo{
			Timestamp: timestamp,
			Category: black_kiwi_db_utils.StringToCategory(category),
			MinRank: rank,
			Coord: black_kiwi_data_structs.Coordinates{
				Longitude: tmpLon,
				Latitude: tmpLat,
			},
			ClusterID: cid,
		}

		tmpList = append(tmpList, reqItem)
	}
	return true, &tmpList
}