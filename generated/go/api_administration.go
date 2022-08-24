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
	"net/http"
)

func GetRequestLocationsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func NewPOIPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
