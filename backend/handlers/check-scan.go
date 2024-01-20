package handlers

import (
	"NewScanner/structs"
  "time"
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

func HandleCheckScan(w http.ResponseWriter, r *http.Request, db *structs.Database, scanData *structs.ScanData, CurrentSignouts *[]structs.Resident) {
      // Handle OPTIONS request for CORS preflight
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method == http.MethodPost {
        var scan structs.RequestData;

        // unmarshals data into scan struct
        getPostDataErr := scan.GetPostData(r.Body);
        if !getPostDataErr.IsNil() {
          response := structs.ScanResponse {
            Success: false,
            Type: "Error",
            Action: "GetPostData",
            Error: getPostDataErr,
          }

          encodeErr := json.NewEncoder(w).Encode(response);
          if encodeErr != nil {
            http.Error(w, "Failed to encode json GetPostData", http.StatusInternalServerError);
          }
          return;
        }

        if strings.ToUpper(scan.Scan) == BREAK {
          scanData.ResetCount();
          response := structs.ScanResponse {
            Success: true,
            Type: "BREAK",
            Action: "BREAK",
          }

          encodeErr := json.NewEncoder(w).Encode(response);
          if encodeErr!= nil {
            http.Error(w, "Faile to encode json Scan == BREAK", http.StatusInternalServerError);
          }
          return;
        }

        if scanData.Count == 0 {
          foundResident, handleResReqErr := handleResidentRequest(w, db, scan, CurrentSignouts);
          if !handleResReqErr.IsNil() {
            // http.Error(w, "Error: The expected request type of Resident ID was not handled properly. Please validate scan and retry. If scan was not of type Resident, then the error is likely due to the fact that a Resident ID was expected", http.StatusInternalServerError);
            responseData := structs.ScanResponse {
              Success: false,
              Action: "handleResidentRequest",
              Type: "Error",
              Error: handleResReqErr,
            }

            encodeErr := json.NewEncoder(w).Encode(responseData);
            if encodeErr != nil {
              http.Error(w, "Failed to encode response data", http.StatusInternalServerError);
            }

            return;
          }

          scanData.Resident = &foundResident;
          scanData.Increment();

        } else {
          handleDevReqErr := handleDeviceRequest(w, db, scan, *scanData, CurrentSignouts);
          if !handleDevReqErr.IsNil() {
            response := structs.ScanResponse {
              Success: false,
              Action: "handleDeviceRequest",
              Error: handleDevReqErr,
            }

            encodeErr := json.NewEncoder(w).Encode(response);
            if encodeErr != nil {
              http.Error(w, "Failed to encode response in handleDeviceRequest count != 0", http.StatusInternalServerError);
            }
            return;
          }
        }

        return;
    }

    http.Error(w, "/api/check-scak Method not POST", http.StatusMethodNotAllowed);
    fmt.Println("Request type not POST");
    return;
}

func handleResidentRequest(w http.ResponseWriter, db *structs.Database, scan structs.RequestData, CurrentSignouts *[]structs.Resident) (structs.Resident, structs.Error) {
  var resident structs.Resident;

  mdoc, convErr := strconv.Atoi(scan.Scan);
  if convErr != nil {
    return resident, structs.Error{
      Place: "check-scan.go handleResidentRequest convErr",
      Message: convErr.Error(),
    };
  }

  foundResident, findResErr := db.FindResident(mdoc);
  if findResErr != nil {
    return resident, structs.Error{
      Place: "check-scan.go handleResidentRequest findResErr",
      Message: findResErr.Error(),
    };
  }

  resident = foundResident;

  responseData := structs.ScanResponse {
    Success: true,
    Type: "Resident",
    Object: resident,
  }

  encodeErr := json.NewEncoder(w).Encode(responseData);
  if encodeErr != nil {
    http.Error(w, "Failed to encode responseData handleResidentRequest", http.StatusInternalServerError);
  }

  return resident, structs.Error{};
}

func handleDeviceRequest(w http.ResponseWriter, db *structs.Database, scan structs.RequestData, scanData structs.ScanData, CurrentSignouts *[]structs.Resident) structs.Error {
  formattedTime := time.Now().Format("01/02/06 15:04:05");

  if len(scan.Scan) < 4 {
    return structs.Error {
      Place: "check-scan.go handleDeviceRequest len < 4",
      Message: "Invalid scan data. Length of scan is less than 4. If the scan is a resident please break then rescan otherwise please verify scan data and retry.",
    }
  }

  prefix := strings.ToUpper(scan.Scan[:3]);
  serial := scan.Scan;
  qrData := scan.Scan[4:];

  switch prefix {
    case COM_SER:
      foundDevice, findDeviceErr := db.FindDevice(COMPUTER, serial);
      if findDeviceErr != nil {
        return structs.Error{
          Place: "check-scan.go handleDeviceRequest COM_SER findDeviceErr",
          Message: findDeviceErr.Error(),
        };
      }

      assign, handleResDeviceErr := scanData.HandleResDevices(db, &foundDevice);
      if !handleResDeviceErr.IsNil() {
        return handleResDeviceErr;
      }

      if assign {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_issued: formattedTime,
        }
        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "ASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }
        
        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest COM_SER assign encodeErr",
            Message: encodeErr.Error(),
          }
        }
      } else {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_returned: formattedTime,
        }

        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "UNASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest COM_SER (!assign) else encodeErr",
            Message: encodeErr.Error(),
          }
        }
      }

    case COM_QR:
      foundDevice, findDeviceErr := db.FindDevice(COMPUTER, qrData);
      if findDeviceErr != nil {
        return structs.Error{
          Place: "check-scan.go handleDeviceRequest COM_QR findDeviceErr",
          Message: findDeviceErr.Error(),
        };
      }

      assign, handleResDeviceErr := scanData.HandleResDevices(db, &foundDevice);
      if !handleResDeviceErr.IsNil() {
        return handleResDeviceErr;
      }

      if assign {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_issued: formattedTime,
        }
        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "ASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest COM_QR assign encodeErr",
            Message: encodeErr.Error(),
          }
        }
      } else {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_returned: formattedTime,
        }

        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }
        response := structs.ScanResponse {
          Success: true,
          Action: "UNASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest COM_QR (!assign) else encodeErr",
            Message: encodeErr.Error(),
          }
        }
      }

    case CAM_QR:
      foundDevice, findDeviceErr := db.FindDevice(CAMERA, qrData);
      if findDeviceErr != nil {
        return structs.Error{
          Place: "check-scan.go handleDeviceRequest CAM_QR findDeviceErr",
          Message: findDeviceErr.Error(),
        };
      }

      assign, handleResDeviceErr := scanData.HandleResDevices(db, &foundDevice);
      if !handleResDeviceErr.IsNil() {
        return handleResDeviceErr;
      }

      if assign {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_issued: formattedTime,
        }
        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "ASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest CAM_QR assign encodeErr",
            Message: encodeErr.Error(),
          }
        }
      } else {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_returned: formattedTime,
        }

        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }
        response := structs.ScanResponse {
          Success: true,
          Action: "UNASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest CAM_SER (!assign) else encodeErr",
            Message: encodeErr.Error(),
          }
        }
      }

    case CAM_SER:
      foundDevice, findDeviceErr := db.FindDevice(CAMERA, serial);
      if findDeviceErr != nil {
        return structs.Error{
          Place: "check-scan.go handleDeviceRequest CAM_SER findDeviceErr",
          Message: findDeviceErr.Error(),
        };
      }

      assign, handleResDeviceErr := scanData.HandleResDevices(db, &foundDevice);
      if !handleResDeviceErr.IsNil() {
        return handleResDeviceErr;
      }

      if assign {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_issued: formattedTime,
        }
        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "ASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest CAM_SER assign encodeErr",
            Message: encodeErr.Error(),
          }
        }
      } else {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_returned: formattedTime,
        }

        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }
        response := structs.ScanResponse {
          Success: true,
          Action: "UNASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest CAM_SER (!assign) else encodeErr",
            Message: encodeErr.Error(),
          }
        }
      }

    case HEA_QR:
      foundDevice, findDeviceErr := db.FindDevice(HEADPHONES, qrData);
      if findDeviceErr != nil {
        return structs.Error{
          Place: "check-scan.go handleDeviceRequest HEA_QR findDeviceErr",
          Message: findDeviceErr.Error(),
        };
      }

      assign, handleResDeviceErr := scanData.HandleResDevices(db, &foundDevice);
      if !handleResDeviceErr.IsNil() {
        return handleResDeviceErr;
      }

      if assign {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_issued: formattedTime,
        }
        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }

        response := structs.ScanResponse {
          Success: true,
          Action: "ASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest HEA_QR assign encodeErr",
            Message: encodeErr.Error(),
          }
        }
      } else {
        newAssignment := structs.Assignment {
          Resident: *scanData.Resident,
          Device: foundDevice,
          Time_returned: formattedTime,
        }

        assignmentErr := db.UpdateAssignmentLog(newAssignment);
        if !assignmentErr.IsNil() {
          return assignmentErr;
        }
        response := structs.ScanResponse {
          Success: true,
          Action: "UNASSIGN",
          Type: "DEVICE",
          Object: foundDevice,
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
          return structs.Error {
            Place: "check-scan.go handleDeviceRequest HEA_QR (!assign) else encodeErr",
            Message: encodeErr.Error(),
          }
        }
      }
  }

  return structs.Error{};
}
