package handlers

import (
	"NewScanner/structs"
	"NewScanner/utils"
	"net/http"
)

const (
  CHECK_SCAN = "api/check-scan"
)

var CurrentSignouts []structs.Resident;

func HandleRoutes(mux *http.ServeMux, db *structs.Database) {
  var scanData ScanData;

  utils.SetCurrentSignOuts(db, &CurrentSignouts);

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handler for the root path
	});

  mux.HandleFunc("/api/check-scan", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
        HandleCheckScan(w, r, db, &scanData);
    }));
}

func setCORSHeaders(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all routes
		w.Header().Set("Access-Control-Allow-Origin", "*");
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS");
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type");

		// Call the next handler in the chain
		next.ServeHTTP(w, r);
	})
}

