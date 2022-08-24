/*
 * Black-Kiwi API
 *
 * API used for the Context Aware System course @ Università degli studi di Bologna a.a. 2021/2022. ITLandfill group 
 *
 * API version: 1.0.0
 * Contact: alessandro.benetton@studio.unibo.it
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type RequestInfo struct {

	Category *Categories `json:"category"`

	Coord *Coordinates `json:"coord"`

	MinRank float64 `json:"minRank,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`
}
