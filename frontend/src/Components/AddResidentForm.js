import { Button, ButtonGroup, FormControl, TextField } from "@mui/material";
import axios from "axios";
import React from "react";

export const AddResidentForm = ({handleCloseDialog, setAddSuccess, setOpenSnackbar}) => {
  const [error, setError] = React.useState(false);
  const [addResidentData, setAddResidentData] = React.useState({
    name: "",
    mdoc: "",
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await axios.post("http://localhost:5432/api/add-new-resident", addResidentData);
      if (response.data.success) {
        setAddSuccess(true);
      } else {
        setAddSuccess(false);
      }
      setOpenSnackbar(true);
      handleCloseDialog();

    } catch (error) {
      console.log(error);
    }
  }

  const handleInputChange = (e) => {
    const {name, value} = e.target;

    if (name === 'mdoc') {
      const isValidInput = /^[0-9]+$/.test(value);
      setError(!isValidInput);
    }

   setAddResidentData({
     ...addResidentData, 
     [name]: value
   })
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <FormControl>
          <TextField name="name" value={addResidentData.name} placeholder="Ex: Samuel Geary" label="Resident Name" helperText="Format: First Last" onChange={handleInputChange} required />
          <TextField name="mdoc" error={error} value={addResidentData.mdoc} label="Resident MDOC" placeholder="Ex: 151110" helperText={error ? "integers only" : ""} onChange={handleInputChange} inputProps={{pattern: "[0-9]+"}} sx={{mt: 1.5}} required />
          <p>* indicates a required field</p>
        <ButtonGroup sx={{alignSelf: "end", mt: 1}}>
          <Button type="submit" variant="contained" sx={{mr: .5}}>Submit</Button>
          <Button onClick={handleCloseDialog} variant="contained" color="error">Close</Button>
        </ButtonGroup>
        </FormControl>
      </form>
    </div>
  )
}

export default AddResidentForm;
