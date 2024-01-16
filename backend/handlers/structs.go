package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"io"
)

type RequestData struct {
  Scan string `json:"scan"`
}

type ScanData struct {
  Resident structs.Resident
  Count int 
}

func (scan *RequestData) getPostData(reqBody io.Reader) error {
  body, readErr := io.ReadAll(reqBody);
  if readErr != nil {
    return readErr;
  }

  unmarshalErr := json.Unmarshal(body, scan);
  if unmarshalErr != nil {
    return unmarshalErr;
  }

  return nil;
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

func (check *ScanData) AssignResident(resident structs.Resident) {
  check.Resident = resident;
}
