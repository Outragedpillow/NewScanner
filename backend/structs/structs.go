package structs

type Resident struct {
  Name string `json:"name"`
  Mdoc int `json:"mdoc"`
  Devices []Device `json:"devices"`
}

type Device struct {
  Type string `json:"type"`
  Serial string `json:"serial"`
  Tag_number int `json:"tag_number"`
  Qr_tag string `json:"qr_tag"`
  Assigned_to *Resident `json:"-"`
}

type Assignment struct {
  Resident Resident `json:"resident"`
  Device Device `json:"device"`
  Time_issued string `json:"time_issued"`
  Time_returned string `json:"time_returned"`
}

type HistoryAssignment struct {
  Mdoc int `json:"mdoc"`
  Name string `json:"name"`
  Type string `json:"type"`
  Serial string `json:"serial"`
  Time_issued string `json:"time_issued"`
  Time_returned string `json:"time_returned"`
}
