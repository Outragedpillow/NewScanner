package handlers

import (
	"NewScanner/structs"
	"NewScanner/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
  COMPUTER = "COMPUTER"
  CAMERA = "CAMERA"
  HEADPHONES = "HEADPHONES"

  COM_SER = "1SR"
  CAM_SER = "214"
  COM_QR = "COM"
  CAM_QR = "CAM"
  HEA_QR = "HEA"

  ASSIGN = "ASSIGN"
  UNASSIGN = "UNASSIGN"
  BREAK = "BREAK"
)

func HandleCheckScan(w http.ResponseWriter, r *http.Request, db *structs.Database, scanData *ScanData) {
      // Handle OPTIONS request for CORS preflight
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method == http.MethodPost {
        var scan RequestData;

        fmt.Println("POST api/check-scan");
        fmt.Println(scanData.Count);

        // unmarshals data into scan struct
        getPostDataErr := scan.getPostData(r.Body);
        if getPostDataErr != nil {
          http.Error(w, "Error: Unable to read request body.", http.StatusInternalServerError);
          fmt.Printf("Place: check-scan.go HandleCheckScan scan.getPostDataErr.\nError: %s", getPostDataErr);
          return;
        }

        if strings.ToUpper(scan.Scan) == BREAK {
          fmt.Println("Before:", scanData.Count)
          scanData.ResetCount();
          fmt.Println("After:", scanData.Count)
          w.Write([]byte("BREAK"));
          return;
        }

        if scanData.Count == 0 {
          foundResident, handleResReqErr := handleResidentRequest(w, db, scan);
          if handleResReqErr != nil {
            http.Error(w, "Error: The expected request type of Resident ID was not handled properly. Please validate scan and retry. If scan was not of type Resident, then the error is likely due to the fact that a Resident ID was expected", http.StatusInternalServerError);
            fmt.Printf("Place: check-scan.go HandleCheckScan Count == 0 handleResidentRequest.\nError: %s", handleResReqErr);
            return;
          }

          scanData.Resident = foundResident;
          scanData.Increment();

        } else {
          handleDevReqErr := handleDeviceRequest(w, db, scan, *scanData);
          if handleDevReqErr != nil {
            http.Error(w, "Error: the expected request type of Device was not handled properly. Please validate scan and retry.", http.StatusInternalServerError);
            fmt.Printf("Place check-scan.go HandleCheckScan Count != 0 handleDeviceRequest.\nError: %s",handleDevReqErr);
            return;
          }
        }

        return;
    }

    http.Error(w, "Method not POST", http.StatusMethodNotAllowed);
    fmt.Println("Not POST");
    return;
}

func handleResidentRequest(w http.ResponseWriter, db *structs.Database, scan RequestData) (structs.Resident, error) {
  var resident structs.Resident;

  mdoc, convErr := strconv.Atoi(scan.Scan);
  if convErr != nil {
    return resident, convErr;
  }

  foundResident, findResErr := db.FindResident(mdoc);
  if findResErr != nil {
    return resident, findResErr;
  }

  resident = foundResident;

  responseData := struct {
    Response structs.Resident
    CurrentSignouts []structs.Resident
  }{
    Response: resident,
    CurrentSignouts: CurrentSignouts,
  }

  response, marshalErr := json.Marshal(responseData);
  if marshalErr != nil {
    return resident, marshalErr;
  }

  w.Header().Set("Content-Type", "application/json");
  w.Write(response); 

  return resident, nil;
}

func handleDeviceRequest(w http.ResponseWriter, db *structs.Database, scan RequestData, scanData ScanData) error {
  if len(scan.Scan) < 4 {
    return fmt.Errorf("Invalid scan data. Length of scan is less than 4. If the scan is a resident please break then rescan otherwise please verify scan data and retry.");
  }

  prefix := strings.ToUpper(scan.Scan[:3]);
  serial := scan.Scan;
  qrData := scan.Scan[4:];

  switch prefix {
    case COM_SER:
      foundDevice, findDeviceErr := db.FindDevice(COMPUTER, serial);
      if findDeviceErr != nil {
        return findDeviceErr;
      }

      assign, handleResDevErr := utils.HandleResDevices(db, &scanData.Resident, &foundDevice, &CurrentSignouts);
      if handleResDevErr != nil {
        return handleResDevErr;
      }

      fmt.Println(qrData, assign);

      // hanlde assingment log
      
    case COM_QR:
      fmt.Println("Hit QR")
      foundDevice, findDeviceErr := db.FindDevice(COMPUTER, qrData);
      if findDeviceErr != nil {
        return findDeviceErr;
      }

      assign, handleResDevErr := utils.HandleResDevices(db, &scanData.Resident, &foundDevice, &CurrentSignouts);
      if handleResDevErr != nil {
        return handleResDevErr;
      }

      fmt.Println(qrData, assign);

      // hanlde assingment log



  }

  return nil;
}
