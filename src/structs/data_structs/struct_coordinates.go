package black_kiwi_data_structs

type Coordinates struct {
	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`
}

var MockCoordinates = []Coordinates{
	{Longitude: 11.320724487304686, Latitude: 44.51254340585983},
	{Longitude: 11.329221725463867, Latitude: 44.48836217722139},
	{Longitude: 11.356172561645508, Latitude: 44.4977297671644},
	{Longitude: 11.34838342666626, Latitude: 44.490398072284904},
	{Longitude: 11.359434127807617, Latitude: 44.49153079516007},
}
