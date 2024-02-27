import { Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, InputLabel, TextField } from "@mui/material";
import React from "react";

const HistoryDialog = ({onSubmit}) => {
  const [openDialog, setOpenDialog] = React.useState(false);
  const [formData, setFormData] = React.useState({
    date: "",
    mdoc: "",
    serial: "",
  })

  const handleSubmit = (event) => {
    event.preventDefault();
    onSubmit(formData);
    setFormData({
      date: "", mdoc: "", display_name: "", serial: ""
    });
    handleClose();
  }

  const handleInputChange = (e) => {
   setFormData({
     ...formData, [e.target.name]: e.target.value
   })
  }

  const handleClose = () => {
    setOpenDialog(false);
  }

  const handleOpen = () => {
    setOpenDialog(true);
  }

 return (
    <div>
      <Button variant="contained" onClick={handleOpen} sx={{ml: 1}}>
        Filter History
      </Button>
      <Dialog open={openDialog} onClose={handleClose}>
        <form onSubmit={handleSubmit}>
        <DialogTitle>History Filter</DialogTitle>
        <DialogContent>
            <FormControl>
              <TextField sx={{mt: 1}} value={formData.date} name="date" label="Date" onChange={handleInputChange} placeholder="format: 01/01/24"/>
              <TextField sx={{mt: 1}} value={formData.mdoc} name="mdoc" label="Mdoc" onChange={handleInputChange} placeholder="ex: 151110"/>
              <TextField sx={{mt: 1}} value={formData.serial} name="serial" label="Device Serial" onChange={handleInputChange} placeholder="ex: 1sR90YFFFF" />
              <TextField sx={{mt: 1}} value={formData.display_name} name="display_name" label="Display Name" onChange={handleInputChange} placeholder="ex: COM-149" />
            </FormControl> 
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">Close</Button>
          <Button type="submit" color="primary">Submit</Button>
        </DialogActions>
        </form>
      </Dialog>
    </div>
  );
}

export default HistoryDialog;
