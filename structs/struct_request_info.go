package black_kiwi
import (
	"time"
)

type RequestInfo struct {

	Category *Categories `json:"category"`

	Coord *Coordinates `json:"coord"`

	MinRank float64 `json:"minRank,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`
}
