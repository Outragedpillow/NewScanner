package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func HandleAddNewDevice(w http.ResponseWriter, r *http.Request, db *structs.Database) {
  var postData structs.Device;

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
        Place: "HandleAddNewDevice Method != POST",
        Message: "Request type not POST",
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewDevice Method != POST failed to encode json", http.StatusInternalServerError);
    }

    return;
  }

  postNewDevDataErr := json.NewDecoder(r.Body).Decode(&postData);
  if postNewDevDataErr != nil {
    response := structs.ScanResponse {
      Success: false,
      RefreshCurr: false,
      Action: "ERROR: Add New Deivce",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewDevice postNewDevDataErr",
        Message: postNewDevDataErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewDevice failed to encode json", http.StatusInternalServerError);
    }
  }

  index := strings.Index(postData.Qr_tag, "-");
  if index != -1 {
    tag_number, convErr := strconv.Atoi(postData.Qr_tag[index+1:]);
    if convErr == nil {
      postData.Tag_number = tag_number;
    }
  } else {
    response := structs.ScanResponse {
      Success: false,
      Type: "ERROR",
      Action: "Convert Qr",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewDevice strings.Index",
        Message: "Error: strings.Index returned -1. Qr_tag does not contain a hyphen",
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewDevice strings.Index else failed to encode json", http.StatusInternalServerError);
    }
  }

  sqlStatment, prepErr := db.Conn.Prepare("insert into devices (type, serial, tag_number, qr_tag) values (?, ?, ?, ?)");
  if prepErr != nil {
    response := structs.ScanResponse {
      Success: false,
      RefreshCurr: false,
      Action: "ERROR: Prepare",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewDevice Prepare",
        Message: prepErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewDevice Prepare failed to encode json", http.StatusInternalServerError);
    }
  }

  _, execErr := sqlStatment.Exec(postData.Type, postData.Serial, postData.Tag_number, strings.ToUpper(postData.Qr_tag));
  if execErr != nil {
    response := structs.ScanResponse {
      Success: false,
      RefreshCurr: false,
      Action: "ERROR: Exec",
      Error: structs.Error{
        Place: "new-device.go HandleAddNewDevice Exec",
        Message: execErr.Error(),
      },
    }

    encodeErr := json.NewEncoder(w).Encode(response);
    if encodeErr != nil {
      http.Error(w, "HandleAddNewDevice Exec failed to encode json", http.StatusInternalServerError);
    }
  }

  response := structs.ScanResponse {
    Success: true,
    Type: "ADD",
    Action: "Added Device",
  }

  encodeErr := json.NewEncoder(w).Encode(response);
  if encodeErr != nil {
    http.Error(w, "HandleAddNewDevice failed to encode json success response", http.StatusInternalServerError);
  }

}
