package structs

import (
	"database/sql"
	"errors"
  "time"
	"fmt"
	"strconv"
	"strings"
)

type Database struct {
  Conn *sql.DB
}

const (
  CAM = "CAMERA"
  COM = "COMPUTER"
  HEA = "HEADPHONES"

  ADD = "ADD"
  REMOVE = "REMOVE"
)

func (database *Database) Open(fileName string) error {
  db, openDbErr := sql.Open("sqlite3", fileName);
  if openDbErr != nil {
    return fmt.Errorf("Failed to open SQLite file with name %s. Error: %w", fileName, openDbErr);
  }

  database.Conn = db;
  return nil;
}

func (database *Database) Close() error {
  if database.Conn != nil {
    return database.Conn.Close();
  }

  return nil;
}

func (database Database) FindResident(mdoc int) (Resident, error) {
    var resident Resident
    var exists int;

    stmnt, pErr := database.Conn.Prepare(`SELECT EXISTS (SELECT 1 FROM residents WHERE mdoc = ? LIMIT 1)`)
    if pErr != nil {
      return resident, pErr;
    }

    defer stmnt.Close();

    qErr := stmnt.QueryRow(mdoc).Scan(&exists);
    if qErr != nil {
      return resident, qErr;
    }

    if exists == 1 {
      sqlStatement, prepErr := database.Conn.Prepare(`
          SELECT residents.name, residents.mdoc, devices.type, devices.serial, devices.tag_number, devices.assigned_to
          FROM residents
          LEFT JOIN devices ON residents.mdoc = devices.assigned_to
          WHERE residents.mdoc = ?
      `)
      if prepErr != nil {
          return resident, prepErr;
      }

      defer sqlStatement.Close()

      rows, queryErr := sqlStatement.Query(mdoc);
      if queryErr != nil {
          return resident, queryErr;
      }
      defer rows.Close();

      for rows.Next() {
          var device Device;
          var deviceType, deviceSerial sql.NullString;
          var deviceTagNumber sql.NullInt32;
          var assignedTo sql.NullInt32;
          err := rows.Scan(&resident.Name, &resident.Mdoc, &deviceType, &deviceSerial, &deviceTagNumber, &assignedTo);
          if err != nil {
              return resident, fmt.Errorf("Error: %w", err);
          }

          // Check if all fields are not NULL
          if deviceType.Valid && deviceSerial.Valid && deviceTagNumber.Valid && assignedTo.Valid {
              device.Type = deviceType.String
              device.Serial = deviceSerial.String
              device.Tag_number = int(deviceTagNumber.Int32)

              // Check if Assigned_to is not NULL
              if assignedTo.Valid {
                  device.Assigned_to = &Resident{Mdoc: int(assignedTo.Int32)}
              }

              resident.Devices = append(resident.Devices, device)
          }
      }

      return resident, nil;

    } else {
      return resident, errors.New("No resident was found");
    }
}

func (database Database) FindDevice(devType string, serial string) (Device, error) {
  deviceType := strings.ToLower(devType);

  var device Device;
  var assignedto sql.NullInt32;
  // use this to see is QR type int rather than serial string
  testIfTag, convErr := strconv.Atoi(serial);
  if convErr != nil {
    sqlStatement, prepErr := database.Conn.Prepare("select type, serial, tag_number, assigned_to from devices where type = ? and serial = ?");
    if prepErr != nil {
      return device, prepErr;
    }

    defer sqlStatement.Close();

    queryErr := sqlStatement.QueryRow(deviceType, serial).Scan(&device.Type, &device.Serial, &device.Tag_number, &assignedto);
    if queryErr != nil {
      if queryErr == sql.ErrNoRows {
        return device, fmt.Errorf("Device %v was not found. Error: %w", device.Tag_number, queryErr);
      }
      return device, queryErr;
    }

    // check if assignedto is valid (not null)
    if assignedto.Valid {
      // create a new resident instance and assign it to device.assigned_to
      device.Assigned_to = &Resident{Mdoc: int(assignedto.Int32)};
    }

    return device, nil;
  }

  sqlStatement, preperr := database.Conn.Prepare("select type, serial, tag_number, assigned_to from devices where type = ? and tag_number = ?");
  if preperr != nil {
    return device, preperr;
  }
  
  defer sqlStatement.Close();

  queryErr := sqlStatement.QueryRow(deviceType, testIfTag).Scan(&device.Type, &device.Serial, &device.Tag_number, &assignedto)
  if queryErr != nil {
    if queryErr == sql.ErrNoRows {
      return device, fmt.Errorf("Device not found. Error: %w", queryErr);
    }
    return device, queryErr;
  }

  // check if assignedto is valid (not null)
  if assignedto.Valid {
    // create a new resident instance and assign it to device.assigned_to
    device.Assigned_to = &Resident{Mdoc: int(assignedto.Int32)};
  }

  return device, nil
}

func (database *Database) UpdateDevice(assignmentState string, deviceType string, serial string, residentMdoc int) error {
  switch assignmentState {
  case "ASSIGN":
    fmt.Println("Assign");
    sqlStatment, prepErr := database.Conn.Prepare("UPDATE devices SET assigned_to = ? WHERE type = ? AND serial = ?");
    if prepErr != nil {
      return prepErr;
    }

    defer sqlStatment.Close();

    _, execErr := sqlStatment.Exec(residentMdoc, deviceType, serial);
    if execErr != nil {
      return execErr;
    }
  case "UNASSIGN":
    fmt.Println("Unassign");
    sqlStatment, prepErr := database.Conn.Prepare("UPDATE devices SET assigned_to = NULL WHERE type = ? AND serial = ?");
    if prepErr != nil {
      return prepErr;
    }

    defer sqlStatment.Close();

    _, execErr := sqlStatment.Exec(deviceType, serial);
    if execErr != nil {
      return execErr;
    }
  }

  return nil;
}

func (db *Database) UpdateCurrentSignOuts(action string, resident *Resident) error {
  switch action {
  case ADD:
    sqlStatment, prepErr := db.Conn.Prepare("INSERT INTO currentsignouts (resident_mdoc, resident_name) VALUES (?, ?)");
    if prepErr != nil {
      return prepErr;
    }

    _, execErr := sqlStatment.Exec(resident.Mdoc, resident.Name);
    if execErr != nil {
      return execErr;
    }

    return nil;

  case REMOVE:
    sqlStatment, prepErr := db.Conn.Prepare("DELETE FROM currentsignouts where resident_mdoc = ?");
    if prepErr != nil {
      return prepErr;
    }

    _, execErr := sqlStatment.Exec(resident.Mdoc);
    if execErr != nil {
      return execErr;
    }

  default:
    return fmt.Errorf("ERROR: action must be either ADD or Remove.");

  }

  return nil;
}

func (database *Database) UpdateAssignmentLog(assingment Assignment) Error {
  formattedDay := time.Now().Format("01/02/06");

  sqlStatement, prepErr := database.Conn.Prepare("INSERT INTO assignments (resident_mdoc, resident_name, device_type, device_serial, time_issued, time_returned, day) VALUES (?, ?, ?, ?, ?, ?, ?)");
  if prepErr != nil {
    return Error {
      Place: "db.go UpdateAssignmentLog Prepare.",
      Message: prepErr.Error(),
    }
  }
  
  defer sqlStatement.Close();

  _, execErr := sqlStatement.Exec(assingment.Resident.Mdoc, assingment.Resident.Name, assingment.Device.Type, assingment.Device.Serial, assingment.Time_issued, assingment.Time_returned, formattedDay);
  if execErr != nil {
    return Error {
      Place: "db.gp UpdateAssignmentLog Exec",
      Message: execErr.Error(),
    }
  }

  return Error{};
}
