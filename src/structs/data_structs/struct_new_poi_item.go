package black_kiwi_data_structs

type NewPoiItem struct {
	Name string `json:"name"`

	Category Categories `json:"category"`

	Rank float32 `json:"rank"`

	Coord Coordinates `json:"coord"`
}
