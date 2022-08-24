package black_kiwi

type Coordinates struct {

	Latitude string `json:"latitude"`

	Longitude string `json:"longitude"`
}

var MockCoordinates = []Coordinates{
	{Latitude: "11.320724487304686", Longitude: "44.51254340585983"},
	{Latitude: "11.329221725463867", Longitude: "44.48836217722139"},
	{Latitude: "11.356172561645508", Longitude: "44.4977297671644"},
	{Latitude: "11.34838342666626", Longitude: "44.490398072284904"},
	{Latitude: "11.359434127807617", Longitude: "44.49153079516007"},

}