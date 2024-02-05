package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"net/http"
)

func HandleGetCurrentSignouts(w http.ResponseWriter, r *http.Request, db *structs.Database, scanData structs.ScanData) {
  // Handle OPTIONS request for CORS preflight
  if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusOK)
      return
  }

  if r.Method == http.MethodGet {
    response := structs.ScanResponse {
      Success: true,
      Type: "CurrentSignout",
      Object: scanData.CurrentSignouts,
      CurrentSignouts: scanData.CurrentSignouts,
      Error: structs.Error{},
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "Failed to encode json Method == Get HandleGetCurrentSignouts", http.StatusInternalServerError);
      return;
    }
      
    return; 
    }

  response := structs.ScanResponse {
    Success: false,
    Type: "ERROR",
    Error: structs.Error{
      Place: "currentsignouts.go HandleGetCurrentSignouts Method != GET",
      Message: "Invalid Request type. Request to this endpoint must be GET",
    },
  }

  encodeErr := json.NewEncoder(w).Encode(response);
  if encodeErr != nil {
    http.Error(w, "Failed to encode json Method != GET", http.StatusInternalServerError);
    return;
  }
}
