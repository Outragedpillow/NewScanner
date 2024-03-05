package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"net/http"
	"strconv"
)

type TempResident struct {
  Name string `json:"name"`
  TempMdoc string `json:"mdoc"`
  Mdoc int
}

func HandleAddNewResident(w http.ResponseWriter, r *http.Request, db *structs.Database) {
  var temp TempResident;
    
  // Handle OPTIONS request for CORS preflight
  if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusOK)
      return
  }

  if r.Method != http.MethodPost {
    response := structs.ScanResponse {
      Success: false,
      Type: "ERROR",
      Error: structs.Error{
        Place: "new-resident.go HandleAddNewResident Method != POST",
        Message: "Request type not POST",
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewResident Method != POST failed to encode json", http.StatusInternalServerError);
    }

    return;
  }

  postNewResDataErr := json.NewDecoder(r.Body).Decode(&temp);
  if postNewResDataErr != nil {
    response := structs.ScanResponse {
      Success: false,
      Type: "ERROR",
      RefreshCurr: false,
      Action: "ERROR: Add New Resident",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewResident postNewResDataErr Decode",
        Message: postNewResDataErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewResident failed to encode json", http.StatusInternalServerError);
    }
  }

  mdoc, convErr := strconv.Atoi(temp.TempMdoc);
  if convErr != nil {
    response := structs.ScanResponse {
      Success: false,
      RefreshCurr: false,
      Type: "ERROR",
      Action: "ERROR: Convert Mdoc",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewResident strconv",
        Message: convErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewResident Exec failed to encode json", http.StatusInternalServerError);
    }
    return;
  }

  temp.Mdoc = mdoc;

  sqlStatment, prepErr := db.Conn.Prepare("insert into residents (name, mdoc) values (?, ?)");
  if prepErr != nil {
    response := structs.ScanResponse {
      Success: false,
      Type: "ERROR",
      RefreshCurr: false,
      Action: "ERROR: Prepare",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewResident Prepare",
        Message: prepErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewResident Prepare failed to encode json", http.StatusInternalServerError);
    }
  }

  _, execErr := sqlStatment.Exec(temp.Name, temp.Mdoc);
  if execErr != nil {
    response := structs.ScanResponse {
      Success: false,
      Type: "ERROR",
      RefreshCurr: false,
      Action: "ERROR: Exec",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewResident Exec",
        Message: execErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewResident Exec failed to encode json", http.StatusInternalServerError);
    }
  }

  response := structs.ScanResponse {
    Success: true,
    Type: "ADD",
    Action: "Added Resident",
  }

  encodeErr := json.NewEncoder(w).Encode(response);
  if encodeErr != nil {
    http.Error(w, "HandleAddNewResident failed to encode json success response", http.StatusInternalServerError);
  }
  
}
