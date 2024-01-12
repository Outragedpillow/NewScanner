package structs

type Resident struct {
  Name string
  Mdoc int
  Devices []Device
}

type Device struct {
  Type string
  Serial string
  Tag_number int
  Assigned_to *Resident
}

