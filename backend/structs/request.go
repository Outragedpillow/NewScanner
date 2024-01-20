package structs

import (
	"encoding/json"
	"errors"
	"io"
)

const (
  ASSIGN = "ASSIGN"
  UNASSIGN = "UNASSIGN"
  )

type ScanResponse struct {
  Success bool `json:"success"`
  Type string `json:"type"`
  Action string `json:"action"`
  Object interface{} `json:"object"`
  Error Error `json:"error"`
}

type Error struct {
  Place string `json:"place"`
  Message string `json:"message"`
}

type RequestData struct {
  Scan string `json:"scan"`
}

type ScanData struct {
  Resident *Resident
  Count int 
  CurrentSignouts *[]Resident
}

func (err Error) IsNil() bool {
  if err.Place == "" && err.Message == "" {
    return true;
  }
  return false;
}

func (scan *RequestData) GetPostData(reqBody io.Reader) Error {
  body, readErr := io.ReadAll(reqBody);
  if readErr != nil {
    return Error{
      Place: "request.go GetPostData readErr",
      Message: readErr.Error(),
    };
  }

  unmarshalErr := json.Unmarshal(body, scan);
  if unmarshalErr != nil {
    return Error {
      Place: "request.go GetPostData unmarshalErr",
      Message: unmarshalErr.Error(),
    };
  }

  return Error{};
}

func (check *ScanData) ResetCount() {
  check.Count = 0;
}

func (check *ScanData) Increment() {
  check.Count += 1;
}

func (check ScanData) GetCount() int {
  return check.Count;
}

func (check *ScanData) AssignResident(resident Resident) {
  check.Resident = &resident;
}

func (scanData *ScanData) HandleResDevices(db *Database, device *Device) (bool, Error) {
  if scanData.CurrentSignouts == nil {
    Error := Error {
      Place: "request.go HandleResDeices CurrentSignouts == nil",
      Message: "CurrentSignouts was not properly initialized before HandleResDevices",
    }
    return false, Error;
  }

  resDevices := append([]Device{}, scanData.Resident.Devices...)
  hasType := false;
  var index int;


  for i, v := range scanData.Resident.Devices {
    if v.Type == device.Type {
      hasType = true;
      index = i;
      break;
    }
  }

  if hasType {
    if scanData.Resident.Devices[index].Assigned_to != nil && device.Assigned_to != nil && scanData.Resident.Devices[index].Serial == device.Serial && scanData.Resident.Devices[index].Assigned_to.Mdoc == device.Assigned_to.Mdoc {
      updateDeviceDbErr := db.UpdateDevice(UNASSIGN, device.Type, device.Serial, scanData.Resident.Mdoc);
      if updateDeviceDbErr != nil {
        Error := Error {
          Place: "request.go HandleResDevices hasType updateDeviceDbErr",
          Message: updateDeviceDbErr.Error(),
        }
        return false, Error;
      }

      resDevices = append(resDevices[:index], resDevices[index+1:]...);
      scanData.Resident.Devices = resDevices;

      device.Assigned_to = nil;

      handleCurrentSignoutsErr := scanData.HandleCurrentSignouts(db, UNASSIGN, scanData.CurrentSignouts, scanData.Resident);
      if handleCurrentSignoutsErr != nil {
        Error := Error {
          Place: "request.go HandleResDevices hasType HandleCurrentSignouts",
          Message: handleCurrentSignoutsErr.Error(),
        }
        return false, Error;
      } 

    } else {
      Error := Error {
        Place: "request.go HandleResDevices hasType else",
        Message: errors.New("Resident already has a device of type " + device.Type + ". Cannot assign more than one of each type to a single resident.").Error(),
      }
      return false, Error;
    }
  } else {
    if device.Assigned_to == nil {
      updateDeviceDbErr := db.UpdateDevice(ASSIGN, device.Type, device.Serial, scanData.Resident.Mdoc);
      if updateDeviceDbErr != nil {
        Error := Error {
          Place: "request.go HandleResDevices Assigned_to == nil (!hasType) else updateDeviceDbErr",
          Message: updateDeviceDbErr.Error(),
        }
        return false, Error;
      }

      device.Assigned_to = scanData.Resident;

      resDevices = append(resDevices, *device);
      scanData.Resident.Devices = resDevices;

      handleCurrentSignoutsErr := scanData.HandleCurrentSignouts(db, ASSIGN, scanData.CurrentSignouts, scanData.Resident);
      if handleCurrentSignoutsErr != nil {
        Error := Error {
          Place: "request.go HandleResDevices Assigned_to == nil (!hasType) else HandleCurrentSignouts",
          Message: handleCurrentSignoutsErr.Error(),
        }
        return false, Error;
      }

      return true, Error{};
    } else {
      Error := Error {
        Place: "request.go HandleResDevices (Assigned_to != nil) else",
        Message: errors.New("Device " + device.Serial + " is assigned to another resident and must be signed back in before reassignment.").Error(),
      }
      return false, Error;
    }
  } 
  return false, Error{};
}

func (scanData *ScanData) HandleCurrentSignouts(db *Database, action string, currentSignouts *[]Resident, resident *Resident) error {
  if scanData.Resident == nil || scanData.CurrentSignouts == nil {
    return errors.New("NIL 2");
  }

  var index int;
  remove := false;
  add := false;
  found := false;

  for i, v := range *currentSignouts {
    if v.Mdoc == resident.Mdoc {
      found = true;
      index = i;
      (*currentSignouts)[i] = *resident;
    }
  }

  if found {
    if action == UNASSIGN {
      if len(resident.Devices) == 0 {
        remove = true;
      }
    }
  } else {
    if action == ASSIGN {
      if len(resident.Devices) > 0 {
        add = true;
      }
    }
  }

  if add && !found {
    *currentSignouts = append(*currentSignouts, *resident);
    updateErr := db.UpdateCurrentSignOuts("ADD", resident);
    if updateErr != nil {
      return updateErr;
    }
  } else if remove && found {
    *currentSignouts = append((*currentSignouts)[:index], (*currentSignouts)[index+1:]...);
    updateErr := db.UpdateCurrentSignOuts("REMOVE", resident);
    if updateErr != nil {
      return updateErr;
    }
  }

  return nil;
}
