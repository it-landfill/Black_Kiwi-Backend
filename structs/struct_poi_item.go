package black_kiwi

type PoiItem struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	Category Categories `json:"category"`

	Rank float64 `json:"rank"`

	Coord *Coordinates `json:"coord"`
}

var MockPOIS = []PoiItem{
	{Id: 0, Name: "POI0", Category: DEPARTMENT, Rank: 5, Coord: &MockCoordinates[0]},
	{Id: 1, Name: "POI1", Category: MUSEUM, Rank: 8, Coord: &MockCoordinates[1]},
	{Id: 2, Name: "POI2", Category: PARK, Rank: 7, Coord: &MockCoordinates[2]},
	{Id: 3, Name: "POI3", Category: THEATER, Rank: 6, Coord: &MockCoordinates[3]},
	{Id: 4, Name: "POI4", Category: HISTORICAL_BUILDING, Rank: 1, Coord: &MockCoordinates[4]},
}