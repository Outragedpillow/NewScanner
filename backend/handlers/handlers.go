package handlers

import (
	"NewScanner/structs"
	"NewScanner/utils"
	"net/http"
)

const (
  CHECK_SCAN = "api/check-scan"
)


func HandleRoutes(mux *http.ServeMux, db *structs.Database) {
  var scanData structs.ScanData = structs.ScanData{
    CurrentSignouts: &[]structs.Resident{},
  }

  utils.SetCurrentSignOuts(db, scanData.CurrentSignouts);

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handler for the root path
	});

  mux.HandleFunc("/api/currentsignouts", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
    HandleGetCurrentSignouts(w, r, db, scanData);
  }))

  mux.HandleFunc("/api/check-scan", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
    HandleCheckScan(w, r, db, &scanData, scanData.CurrentSignouts);
  }));

  mux.HandleFunc("/api/history", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
    HandleGetHistory(w, r, db);
  }));

  mux.HandleFunc("/api/add-new-resident", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
    HandleAddNewResident(w, r, db);
  }));

  mux.HandleFunc("/api/add-new-device", setCORSHeaders(func(w http.ResponseWriter, r *http.Request) {
    HandleAddNewDevice(w, r, db);
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

