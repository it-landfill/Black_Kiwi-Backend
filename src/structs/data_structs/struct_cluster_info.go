package black_kiwi_data_structs

import (
	"time"
)

type ClusterInfo struct {
	Category Categories `json:"category, omitempty"`

	Coord Coordinates `json:"coord"`

	MinRank float64 `json:"minRank,omitempty"`

	Timestamp time.Time `json:"timestamp"`

	ClusterID int `json:"clusterID"`
}

var MockClusterInfo = []ClusterInfo{
	{Category: DEPARTMENT, Coord: MockCoordinates[0], MinRank: 2, Timestamp: time.Now(), ClusterID: 0},
	{Category: MUSEUM, Coord: MockCoordinates[0], MinRank: 3, Timestamp: time.Now(), ClusterID: 0},
	{Category: PARK, Coord: MockCoordinates[0], MinRank: 4, Timestamp: time.Now(), ClusterID: 1},
	{Category: HISTORICAL_BUILDING, Coord: MockCoordinates[2], MinRank: 5, Timestamp: time.Now(), ClusterID: 2},
	{Category: THEATER, Coord: MockCoordinates[0], MinRank: 6, Timestamp: time.Now(), ClusterID: 3},
	{Category: PARK, Coord: MockCoordinates[0], MinRank: 7, Timestamp: time.Now(), ClusterID: 3},
}
