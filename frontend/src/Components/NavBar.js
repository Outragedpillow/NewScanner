import { NavLink } from 'react-router-dom';
import React from 'react';
import HistorySharpIcon from '@mui/icons-material/HistorySharp';
import CoPresentSharpIcon from '@mui/icons-material/CoPresentSharp';
import InfoSharpIcon from '@mui/icons-material/InfoSharp';
import { Alert, Dialog, DialogContent, DialogTitle, FormControl, FormControlLabel, IconButton, Radio, RadioGroup, Snackbar, Tooltip, Typography } from '@mui/material';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import AddResidentForm from './AddResidentForm';
import AddDeviceForm from './AddDeviceDialog';

const NavBar = () => {
  const [openAddDialog, setOpenAddDialog] = React.useState(false);
  const [openSnackbar, setOpenSnackbar] = React.useState(false);
  const [radioValue, setRadioValue] = React.useState("resident");
  const [addSuccess, setAddSuccess] = React.useState(false);

  const handleOpenDialog = () => {
    setOpenAddDialog(true);
  }

  const handleCloseDialog = () => {
    setRadioValue("resident")
    setOpenAddDialog(false);
  }

  const handleRadioChange = (e) => {
    setRadioValue(e.target.value);
  }

  return (
    <div>
      <Snackbar open={openSnackbar} onClose={() => {setOpenSnackbar(false)}} autoHideDuration={2000} anchorOrigin={{vertical: 'top', horizontal: 'center'}}  >
         <Alert onClose={() => {setOpenSnackbar(false)}} variant='filled' severity={addSuccess ? 'success' : 'error'}>
           {addSuccess ? "Successfully Added!" : "Add Operation Failed!"}
         </Alert>
      </Snackbar>
      <div>
        <IconButton >
           <NavLink to={"/home/"}>
             <Tooltip title={<Typography sx={{fontSize: 30}}>Current Signouts</Typography>} placement='left' 
               componentsProps={{
                  tooltip: {
                    sx: {
                      bgcolor: '#2196f3',
                      color: 'white'
                    },
                  },
                }}>
              <CoPresentSharpIcon sx={{mt: 3.5, fontSize: '50px', color: 'blue'}}/>
            </Tooltip>
          </NavLink>
        </IconButton>
        <IconButton >
           <NavLink to={"/home/history"}>
             <Tooltip title={<Typography sx={{fontSize: 30}}>History</Typography>} placement='left' 
              componentsProps={{
                tooltip: {
                  sx: {
                    bgcolor: '#2196f3',
                    color: 'white'
                  },
                },
              }}>
              <HistorySharpIcon sx={{mt: 3.5, fontSize: '50px', color: 'blue'}}/>
            </Tooltip>
          </NavLink>
        </IconButton>
        <IconButton onClick={handleOpenDialog}>
          <Tooltip title={<Typography sx={{fontSize: 30}}>Add Resident or Device</Typography>} placement='left' 
            componentsProps={{
              tooltip: {
                sx: {
                  bgcolor: '#2196f3',
                  color: 'white'
                },
              },
            }}>
            <AddCircleOutlineIcon sx={{mt: 3.5, fontSize: '50px', color: 'blue'}}/>
          </Tooltip>
        </IconButton>
        <IconButton >
           <NavLink to={"/home/about"}>
             <Tooltip title={<Typography sx={{fontSize: 30}}>About Page</Typography>} placement='left' 
                componentsProps={{
                  tooltip: {
                    sx: {
                      bgcolor: '#2196f3',
                      color: 'white'
                    },
                  },
              }}>
               <InfoSharpIcon sx={{mt: 3.5, fontSize: '50px', color: 'blue'}}/>
             </Tooltip>
          </NavLink>
        </IconButton>
      </div>
      <Dialog open={openAddDialog} onClose={handleCloseDialog}>
        <DialogTitle sx={{alignSelf: 'center'}}>{radioValue == "resident" ? "Add New Resident" : "Add New Device"}</DialogTitle>
          <DialogContent>
            <FormControl>
            <RadioGroup defaultValue="resident" row sx={{ml: 3}} onChange={handleRadioChange}>
              <FormControlLabel control={<Radio />} value="resident" label="Resident" />
              <FormControlLabel control={<Radio />} value="device" label="Device" />
            </RadioGroup>  
            </FormControl>
            {radioValue == 'resident' ? <AddResidentForm handleCloseDialog={handleCloseDialog} setAddSuccess={setAddSuccess} addSuccess={addSuccess} setOpenSnackbar={setOpenSnackbar} /> : <AddDeviceForm handleCloseDialog={handleCloseDialog} setAddSuccess={setAddSuccess} setOpenSnackbar={setOpenSnackbar} />}
          </DialogContent>
      </Dialog>
    </div>
  )
};

export default NavBar;
