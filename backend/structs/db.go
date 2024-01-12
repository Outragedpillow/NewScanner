package structs

import (
	"database/sql"
	"errors"
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
)

func (database *Database) Open(fileName string) error {
  db, openDbErr := sql.Open("sqlite3", fileName);
  if openDbErr != nil {
    return fmt.Errorf("db.go Open, failed to open SQLite file with name %s. Error: %w", fileName, openDbErr);
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
      return resident, fmt.Errorf("db.go FindResident pErr. Error: %w", pErr);
    }

    defer stmnt.Close();

    qErr := stmnt.QueryRow(mdoc).Scan(&exists);
    if qErr != nil {
      return resident, fmt.Errorf("db.go FindResident qErr. Error: %w", qErr);
    }

    if exists == 1 {
      sqlStatement, prepErr := database.Conn.Prepare(`
          SELECT residents.name, residents.mdoc, devices.type, devices.serial, devices.tag_number, devices.assigned_to
          FROM residents
          LEFT JOIN devices ON residents.mdoc = devices.assigned_to
          WHERE residents.mdoc = ?
      `)
      if prepErr != nil {
          return resident, fmt.Errorf("db.go FindResident Prepare. Error: %w", prepErr)
      }

      defer sqlStatement.Close()

      rows, queryErr := sqlStatement.Query(mdoc)
      if queryErr != nil {
          return resident, fmt.Errorf("db.go FindResident Query. Error: %w", queryErr)
      }
      defer rows.Close()

      for rows.Next() {
          var device Device
          var deviceType, deviceSerial sql.NullString
          var deviceTagNumber sql.NullInt32
          var assignedTo sql.NullInt32
          err := rows.Scan(&resident.Name, &resident.Mdoc, &deviceType, &deviceSerial, &deviceTagNumber, &assignedTo)
          if err != nil {
              return resident, fmt.Errorf("db.go FindResident Scan. Error: %w", err)
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
  fmt.Println(serial)

  deviceType := strings.ToLower(devType)

  var device Device
  var assignedto sql.NullInt32
  // use this to see is QR type int rather than serial string
  testIfTag, convErr := strconv.Atoi(serial);
  if convErr != nil {
    fmt.Println("Hit inside convErr")
    
    sqlStatement, prepErr := database.Conn.Prepare("select type, serial, tag_number, assigned_to from devices where type = ? and serial = ?");
    if prepErr != nil {
      return device, fmt.Errorf("db.go FindDevice prepErr != nil. Error: %w", prepErr)
    }

    defer sqlStatement.Close();

    queryErr := sqlStatement.QueryRow(deviceType, serial).Scan(&device.Type, &device.Serial, &device.Tag_number, &assignedto);
    if queryErr != nil {
      if queryErr == sql.ErrNoRows {
        return device, fmt.Errorf("Device %v was not found. Error: %w", device.Tag_number, queryErr);
      }
      return device, fmt.Errorf("db.go FindDevice queryRowErr != nil. Error: %w", queryErr);
    }

    // check if assignedto is valid (not null)
    if assignedto.Valid {
      // create a new resident instance and assign it to device.assigned_to
      device.Assigned_to = &Resident{Mdoc: int(assignedto.Int32)}
    }

    return device, nil;
  }

  sqlStatement, preperr := database.Conn.Prepare("select type, serial, tag_number, assigned_to from devices where type = ? and tag_number = ?");
  if preperr != nil {
    return device, fmt.Errorf("db.go findDevice prepare. error: %w", preperr)
  }
  
  defer sqlStatement.Close();

  queryErr := sqlStatement.QueryRow(deviceType, testIfTag).Scan(&device.Type, &device.Serial, &device.Tag_number, &assignedto)
  if queryErr != nil {
    if queryErr == sql.ErrNoRows {
      return device, fmt.Errorf("Device not found. Error: %w", queryErr);
    }
    return device, fmt.Errorf("db.go findDevice queryrow.scan. error: %w", queryErr)
  }

  // check if assignedto is valid (not null)
  if assignedto.Valid {
    // create a new resident instance and assign it to device.assigned_to
    device.Assigned_to = &Resident{Mdoc: int(assignedto.Int32)}
  }

  return device, nil
}

