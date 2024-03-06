import { Button, ButtonGroup, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, TextField } from "@mui/material";
import React from "react";

const HistoryDialog = ({onSubmit}) => {
  const [openFilterDialog, setOpenFilterDialog] = React.useState(false);
  const [dateError, setDateError] = React.useState(false);
  const [mdocError, setMdocError] = React.useState(false);
  const [displayNameError, setDisplayNameError] = React.useState(false);
  const [formData, setFormData] = React.useState({
    date: "",
    mdoc: "",
    serial: "",
    display_name: "",
  })

  const handleSubmit = (event) => {
    event.preventDefault();
    onSubmit(formData);
    setFormData({
      date: "", mdoc: "", display_name: "", serial: "",
    });
    handleFilterClose();
  }

  const handleInputChange = (e) => {
    const {name, value} = e.target;

    if (name === "date") {
      const isValidInput = /^[0-9]{2}\/[0-9]{2}\/[0-9]{2}$/.test(value);
      setDateError(!isValidInput);
    }

    if (name === 'mdoc') {
      const isValidInput = /^[0-9]+$/.test(value);
      setMdocError(!isValidInput);
    }

    if (name === 'display_name') {
      const isValidInput = /^[A-Za-z0-9]+-[A-Za-z0-9]+$/.test(value);
      setDisplayNameError(!isValidInput);
    }

    setFormData({
      ...formData, [e.target.name]: e.target.value
    })
  }

  const handleFilterClose = () => {
    setFormData({
      date: "", mdoc: "", display_name: "", serial: "",
    });
    setOpenFilterDialog(false)
  }

  const handleFilterOpen = () => {
    setOpenFilterDialog(true);
  }

 return (
    <div>
      <ButtonGroup>
        <Button variant="contained" onClick={handleFilterOpen} sx={{ml: 1}}>
          Filter History
        </Button>
      </ButtonGroup>
      <Dialog open={openFilterDialog} onClose={handleFilterClose}>
        <form onSubmit={handleSubmit}>
        <DialogTitle>History Filter</DialogTitle>
        <DialogContent>
            <FormControl>
              <TextField sx={{mt: 1}} value={formData.date} error={dateError} name="date" label="Date" onChange={handleInputChange} placeholder="Format: 01/01/24"
                helperText={dateError ? "required format: mm/dd/yy" : ""}  
                inputProps={{
                  pattern: "[0-9]{2}/[0-9]{2}/[0-9]{2}"}}
              />
              <TextField sx={{mt: 1}} value={formData.mdoc} error={mdocError} name="mdoc" label="Mdoc" onChange={handleInputChange} placeholder="Format: 151110"
                helperText={mdocError ? "intergers only" : ""}  
                inputProps={{
                  pattern: "[0-9]+"}}
                />
              <TextField sx={{mt: 1}} value={formData.serial} name="serial" label="Device Serial" onChange={handleInputChange} placeholder="Ex: 1sR90YFFFF" />
              <TextField sx={{mt: 1}} value={formData.display_name} error={displayNameError} name="display_name" label="Display Name" onChange={handleInputChange} placeholder="Format: COM-149"
                helperText={displayNameError ? "must contain a hyphen followed by a number" : ""}  
                inputProps={{
                  pattern: "[A-Za-z0-9]+-[A-Za-z0-9]+"}}
              />
            </FormControl> 
        </DialogContent>
        <DialogActions>
          <Button onClick={handleFilterClose} color="primary">Close</Button>
          <Button type="submit" color="primary">Submit</Button>
        </DialogActions>
        </form>
      </Dialog>
    </div>
  );
}

export default HistoryDialog;
