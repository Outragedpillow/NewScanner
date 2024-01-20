package utils

import (
	"NewScanner/structs"
	"errors"
	"fmt"
)

const (
  ASSIGN = "ASSIGN"
  UNASSIGN = "UNASSIGN"
)

// bool returned indicates: true = assign, false = unassign
func HandleResDevices(db *structs.Database, resident *structs.Resident, device *structs.Device, currentSignOuts *[]structs.Resident) (bool, error) {
  hasType := false;
  var index int;

  for i := range resident.Devices {
    if resident.Devices[i].Type == device.Type {
      hasType = true;
      index = i;
      break;
    }  
  }

  if hasType {
    fmt.Println("hasType")
    if resident.Devices[index].Assigned_to != nil && device.Assigned_to != nil && resident.Devices[index].Assigned_to.Mdoc == device.Assigned_to.Mdoc && resident.Devices[index].Serial == device.Serial {
      // update db
      updateDbErr := db.UpdateDevice(UNASSIGN, device.Type, device.Serial, resident.Mdoc);
      if updateDbErr != nil {
        return false, fmt.Errorf("utils.go handleResDevices hasType UpdateDevice. Error %w", updateDbErr);
      }
      // remove from devices slice
      resident.Devices = append(resident.Devices[:index], resident.Devices[index+1:]...);
      // unassign device field
      device.Assigned_to = nil;

      handleCurrentSignOuts(db, currentSignOuts, *resident, UNASSIGN);

      return false, nil;
    } else {
      return false, errors.New("Resident already has a device of type " + device.Type + ". Cannot assign more than one of each type to a single resident.");
    }
  } else {
    fmt.Println("!hasType")
    if device.Assigned_to == nil {
      // update db
      updateDbErr := db.UpdateDevice(ASSIGN, device.Type, device.Serial, resident.Mdoc);
      if updateDbErr != nil {
        return false, fmt.Errorf("utils.go handleResDevices !hasType UpdateDevice. Error %w", updateDbErr);
      }
      // assign device field
      device.Assigned_to = resident;
      // add to devices slice
      resident.Devices = append(resident.Devices, *device);

      handleCurrentSignOuts(db, currentSignOuts, *resident, ASSIGN);

      return true, nil;
    } else {
      return false, errors.New("Device " + device.Serial + " is assigned to another resident and must be signed back in before reassignment.");
    }

  }

}

func handleCurrentSignOuts(db *structs.Database, currentSignOuts *[]structs.Resident, resident structs.Resident, action string) {
  var index int;
  remove := false;
  add := false;
  found := false;

  for i, v := range *currentSignOuts {
    if v.Mdoc == resident.Mdoc {
      found = true;
      index = i;
      (*currentSignOuts)[i] = resident;
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
    *currentSignOuts = append(*currentSignOuts, resident);
    updateErr := db.UpdateCurrentSignOuts("ADD", &resident);
    if updateErr != nil {
      fmt.Println(updateErr);
    }
  } else if remove && found {
    *currentSignOuts = append((*currentSignOuts)[:index], (*currentSignOuts)[index+1:]...);
    updateErr := db.UpdateCurrentSignOuts("REMOVE", &resident);
    if updateErr != nil {
      fmt.Println(updateErr);
    }
  }

}

func SetCurrentSignOuts(db *structs.Database, currentSignOuts *[]structs.Resident) error {
  var mdoc int;

  sqlStatment, prepErr := db.Conn.Prepare("SELECT resident_mdoc FROM currentsignouts");
  if prepErr != nil {
    return fmt.Errorf("utils.go SetCurrentSignOuts Prepare. Error: %w", prepErr);
  }

  rows, queryErr := sqlStatment.Query();
  if queryErr != nil {
    return fmt.Errorf("utils.go SetCurrentSignOuts Query. Error: %w", queryErr);
  }

  for rows.Next() {
    scanErr := rows.Scan(&mdoc);
    if scanErr != nil {
     fmt.Println("Scan Error setting currentSignOuts");
     continue;
    }

    foundResident, findResidentErr := db.FindResident(mdoc);
    if findResidentErr != nil {
      fmt.Println("Invalid resident in currentsignouts table");
      continue;
    }

    *currentSignOuts = append(*currentSignOuts, foundResident);
  }

  return nil;
}

