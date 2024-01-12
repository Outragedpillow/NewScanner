package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type RequestData struct {
  Scan string `json:"scan"`
}

const (
  COMPUTER = "COMPUTER"
  CAMERA = "CAMERA"
  HEADPHONES = "HEADPHONES"

  COM_SER = "1SR"
  CAM_SER = "214"
  COM_QR = "COM"
  CAM_QR = "CAM"
  HEA_QR = "HEA"
)

func HandleCheckScan(w http.ResponseWriter, r *http.Request, db *structs.Database, scanData *ScanData) {
      // Handle OPTIONS request for CORS preflight
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method == http.MethodPost {
        fmt.Println("POST api/check-scan");

        body, bodyErr := io.ReadAll(r.Body);
        if bodyErr != nil {
            http.Error(w, "Failed to read request body", http.StatusBadRequest);
            return;
        }

        var scan RequestData;

        unmarshalErr := json.Unmarshal(body, &scan);
        if unmarshalErr != nil {
            http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest);
            return;
        }

        if scanData.GetCount() == 0 {
          foundResident, findResErr := handleFindResident(db, scan.Scan);
          if findResErr != nil {
            http.Error(w, fmt.Sprintf("%s", findResErr), http.StatusInternalServerError);
            return;
          }
          

          w.Header().Set("Content-Type", "application/json");

          response, marshalErr := json.Marshal(foundResident);
          if marshalErr != nil {
              http.Error(w, "Failed to marshal response JSON", http.StatusInternalServerError);
              return;
          }

          w.Write(response);
          scanData.Increment();
          return;
        } else {
          if len(scan.Scan) > 4 {
            http.Error(w, "Scan length less than four is not being handled yet.", http.StatusInternalServerError);
            return;
          }

          switch strings.ToUpper(scan.Scan[:3]) {
          case COM_QR:
            foundDevice, findDeviceErr := handleFindDevice(db, COMPUTER, scan.Scan[4:]);
            if findDeviceErr != nil {
              http.Error(w, fmt.Sprintf("%s", findDeviceErr), http.StatusInternalServerError);
              return;
            }

            response, marshalErr := json.Marshal(foundDevice);
            if marshalErr != nil {
              http.Error(w, "Failed to Marshal foundDevice in COM_QR", http.StatusInternalServerError);
              return;
            }

            w.Header().Set("Content-Type", "application/json");
            w.Write(response);
            return;
          }
          
          
        }
    }

    http.Error(w, "Method not POST", http.StatusMethodNotAllowed);
    fmt.Println("Not POST");
    return;
}

func handleFindResident(db *structs.Database, mdocStr string) (structs.Resident, error) {
    var resident structs.Resident;

    mdoc, convErr := strconv.Atoi(mdocStr);
    if convErr != nil {
      return resident, fmt.Errorf("check-scan.go handleResidentScan convErr. Error: %w", convErr);
    };

    foundResident, findResErr := db.FindResident(mdoc);
    if findResErr != nil {
      return resident, fmt.Errorf("check-scan.go handleResidentScan FindResident. Error: %w", findResErr);
    };

    resident = foundResident;

    return foundResident, nil;
}

func handleFindDevice(db *structs.Database, devType string, serial string) (structs.Device, error) {
  var device structs.Device;

  foundDevice, findDeviceErr := db.FindDevice(devType, serial);
  if findDeviceErr != nil {
    return device, fmt.Errorf("check-scan.go handleFindDevice FindDevice. Error: %w", findDeviceErr);
  }

  device = foundDevice;

  return device, nil;
}
