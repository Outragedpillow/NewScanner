import { Button, ButtonGroup, Card, FormControl, TextField } from "@mui/material";
import axios from "axios";
import React, { useState } from "react";

export const AddDeviceForm = ({handleCloseDialog, setAddSuccess, setOpenSnackbar}) => {
  const [error, setError] = React.useState(false);
  const [addDeviceData, setAddDeviceData] = React.useState({
    type: "",
    serial: "",
    qr_tag: "",
  })

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    const response = await axios.post("http://localhost:5432/api/add-new-device", addDeviceData);
    if (response.data.success) {
      setAddSuccess(true);
    } else {
      setAddSuccess(false);
    }
    setOpenSnackbar(true);
    handleCloseDialog();
  }

  const handleInputChange = (e) => {
    const { name, value } = e.target;

    if (name === "qr_tag") {
      const isValidInput = /^[A-Za-z0-9]+-[0-9]+$/.test(value);
      setError(!isValidInput);
    }

    setAddDeviceData({
      ...addDeviceData,
      [name]: value
    });
  }

  return (
    <form onSubmit={handleSubmit}>
      <FormControl>
        <TextField name="type" value={addDeviceData.type} label="Device Type" placeholder="Ex: Computer" onChange={handleInputChange} sx={{mt: 0}} required />
        <TextField name="serial" value={addDeviceData.serial} label="Device Serial" placeholder="Ex: R90YFFFF" onChange={handleInputChange} sx={{mt: 2}} required />
        <TextField
          name="qr_tag"
          error={error}
          value={addDeviceData.qr_tag}
          label="Device Qr Tag"
          placeholder="Format: COM-143"
          helperText={error ? "must contain a hyphen followed by a number" : ""}
          onChange={handleInputChange}
          inputProps={{ pattern: "[A-Za-z0-9]+-[0-9]+" }}
          sx={{ mt: 2 }}
          required
        />
        <p>* indicates a required field</p>
        <ButtonGroup sx={{ alignSelf: "end", mt: 1 }}>
          <Button type="submit" variant="contained" sx={{ mr: 0.5 }}>Submit</Button>
          <Button onClick={handleCloseDialog} variant="contained" color="error">Close</Button>
        </ButtonGroup>
      </FormControl>
    </form>
  )
}

export default AddDeviceForm;
