package black_kiwi_structs

type NewPoiItem struct {

	Name string `json:"name"`

	Category *Categories `json:"category"`

	Coord *Coordinates `json:"coord"`
}
