package black_kiwi_structs

type PoiItem struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	Category Categories `json:"category"`

	Rank float32 `json:"rank"`

	Coord *Coordinates `json:"coord"`
}

var MockPOIS = []PoiItem{
	{Id: 0, Name: "POI0", Category: DEPARTMENT, Rank: 5.2, Coord: &MockCoordinates[0]},
	{Id: 1, Name: "POI1", Category: MUSEUM, Rank: 8.0, Coord: &MockCoordinates[1]},
	{Id: 2, Name: "POI2", Category: PARK, Rank: 7.8, Coord: &MockCoordinates[2]},
	{Id: 3, Name: "POI3", Category: THEATER, Rank: 6.4, Coord: &MockCoordinates[3]},
	{Id: 4, Name: "POI4", Category: HISTORICAL_BUILDING, Rank: 1.9, Coord: &MockCoordinates[4]},
}