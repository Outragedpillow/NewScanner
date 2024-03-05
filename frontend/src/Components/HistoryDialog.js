import { Button, ButtonGroup, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, TextField } from "@mui/material";
import React from "react";

const HistoryDialog = ({onSubmit}) => {
  const [openFilterDialog, setOpenFilterDialog] = React.useState(false);
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
              <TextField sx={{mt: 1}} value={formData.date} name="date" label="Date" onChange={handleInputChange} placeholder="Format: 01/01/24"/>
              <TextField sx={{mt: 1}} value={formData.mdoc} name="mdoc" label="Mdoc" onChange={handleInputChange} placeholder="Format: 151110"/>
              <TextField sx={{mt: 1}} value={formData.serial} name="serial" label="Device Serial" onChange={handleInputChange} placeholder="Ex: 1sR90YFFFF" />
              <TextField sx={{mt: 1}} value={formData.display_name} name="display_name" label="Display Name" onChange={handleInputChange} placeholder="Format: COM-149" />
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
