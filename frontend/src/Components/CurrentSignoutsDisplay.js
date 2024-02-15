import { Box, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material";
import { makeStyles } from "@mui/styles";
import React from "react";
import { GetDataContext } from "../Pages/Home";

const useStyles = makeStyles({
  tableCell: {
    padding: '0px 0px 0px 0px',
  },

  tableRow: {
    '& > *': {
      padding: '0px 0px 0px 0px',
    },
  },
});

const CurrentSignoutsDisplay = () => {
  const classes = useStyles();
  const currentSignouts = React.useContext(GetDataContext);

  if (currentSignouts === undefined) {
    return (
      <Box sx={{ height: "95vh", overflow: "auto" }}>
        <Table><TableBody><TableRow><TableCell><Typography variant="h3">No Devices Assigned</Typography></TableCell></TableRow></TableBody></Table>
      </Box>
    );
  } 

  return (
    <Box sx={{ height: "95vh", overflow: "auto", ml: .5 }}>
      <Table sx={{ minWidth: 650 }} size="small">
        <TableHead>
          <TableRow>
            <TableCell sx={{pl: 0}}><Typography fontSize={30}>Name</Typography></TableCell>
            <TableCell sx={{pl: 0}} ><Typography fontSize={30}>Mdoc</Typography></TableCell>
            <TableCell sx={{pl: 0}} ><Typography fontSize={30}>Devices</Typography></TableCell>
          </TableRow>
        </TableHead>
        <br />
        <TableBody>
          {currentSignouts.map((item) => (
            <TableRow key={item.mdoc} className={classes.tableRow}>
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                <Typography fontSize={25} sx={{mr: -10, pr: 0, pl: 0, pr: 0}}>{item.name}</Typography>
              </TableCell>
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                <Typography fontSize={20} sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>{item.mdoc}</Typography>
              </TableCell>
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                {item.devices.map((device) => {
                  return <Typography fontSize={20} key={device.serial}>Type: {device.type}, Tag Number: {device.tag_number}</Typography>;
                })} 
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Box>
  );
};

export default CurrentSignoutsDisplay;
