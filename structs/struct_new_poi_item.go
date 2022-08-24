package black_kiwi

type NewPoiItem struct {

	Name string `json:"name"`

	Category *Categories `json:"category"`

	Coord *Coordinates `json:"coord"`
}
