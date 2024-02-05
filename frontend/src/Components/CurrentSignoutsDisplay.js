import { Box, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material";
import axios from "axios";
import React from "react";

const CurrentSignoutsDisplay = ({ currentSignouts }) => {
  if (currentSignouts === undefined) {
    return (
      <Box sx={{ height: "90vh", overflow: "auto" }}>
        <Table><TableBody><TableRow><TableCell><Typography variant="h3">No Devices Assigned</Typography></TableCell></TableRow></TableBody></Table>
      </Box>
    );
  } 


  return (
    <Box sx={{ height: "95vh", overflow: "auto" }}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell><Typography>Name</Typography></TableCell>
            <TableCell><Typography>Mdoc</Typography></TableCell>
            <TableCell><Typography>Devices</Typography></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {currentSignouts.map((item) => (
            <TableRow key={item.mdoc}>
              <TableCell>
                <Typography>{item.name}</Typography>
              </TableCell>
              <TableCell>
                <Typography>{item.mdoc}</Typography>
              </TableCell>
              <TableCell>
                {item.devices.map((device) => {
                  return <Typography key={device.serial}>Type: {device.type}, Tag Number: {device.tag_number}</Typography>;
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
