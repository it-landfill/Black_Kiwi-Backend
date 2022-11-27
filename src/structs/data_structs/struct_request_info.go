package black_kiwi_data_structs

import (
	"time"
)

type RequestInfo struct {
	Category Categories `json:"category"`

	Coord Coordinates `json:"coord"`

	MinRank float64 `json:"minRank,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`
}

var MockRequestInfo = []RequestInfo{
	{Category: DEPARTMENT, Coord: MockCoordinates[0], MinRank: 2, Timestamp: time.Now()},
	{Category: MUSEUM, Coord: MockCoordinates[0], MinRank: 3, Timestamp: time.Now()},
	{Category: PARK, Coord: MockCoordinates[0], MinRank: 4, Timestamp: time.Now()},
	{Category: HISTORICAL_BUILDING, Coord: MockCoordinates[2], MinRank: 5, Timestamp: time.Now()},
	{Category: THEATER, Coord: MockCoordinates[0], MinRank: 6, Timestamp: time.Now()},
	{Category: PARK, Coord: MockCoordinates[0], MinRank: 7, Timestamp: time.Now()},
}
