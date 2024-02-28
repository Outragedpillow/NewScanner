import { Box, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material";
import React from "react";
import { GetDataContext } from "../Pages/Home";

const CurrentSignoutsDisplay = () => {
  const currentSignouts = React.useContext(GetDataContext);

  if (currentSignouts === undefined) {
    return (
      <Box sx={{ height: "95vh", overflow: "auto" }}>
        <Table><TableBody><TableRow><TableCell><Typography variant="h3">No Devices Assigned</Typography></TableCell></TableRow></TableBody></Table>
      </Box>
    );
  } 

  let arrayLength = currentSignouts.length;

  return (
    <Box sx={{ height: "95vh", overflow: "auto", ml: .5 }}>
      <Table stickyHeader sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
        <TableHead>
          <TableRow>
            <TableCell sx={{pl: 0}}><Typography fontSize={30}>Name</Typography></TableCell>
            <TableCell sx={{pl: 0}} ><Typography fontSize={30}>Mdoc</Typography></TableCell>
            <TableCell sx={{pl: 0}} ><Typography fontSize={30}>Devices</Typography></TableCell>
          </TableRow>
        </TableHead>
        <br />
        <TableBody>
          {currentSignouts.map((item, index) => (
            <TableRow key={index}>
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                <Typography fontSize={20} sx={{mr: -10}}>{item.name}</Typography>
              </TableCell>
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                <Typography fontSize={20} sx={{pt: -2}}>{item.mdoc}</Typography>
              </TableCell>
              {arrayLength <= 8 ? 
              <TableCell sx={{mr: 0, pr: 0, pl: 0, pr: 0}}>
                {item.devices.map((device, index2) => {
                  return <Typography fontSize={20} key={index2}>Type: {device.type}, Tag Number: {device.tag_number}</Typography>;
                })} 
              </TableCell>
              : 
              <TableCell sx={{display: "flex", ml: 0, pl: 0}}>
                {item.devices.map((device, index3) => {
                  return <Typography fontSize={20} key={index3} sx={{pr: 0, mt: 0}}>{index3 < 2 && index3 < item.devices.length - 1 ? device.qr_tag + ","  : device.qr_tag}&nbsp;</Typography>;
                })} 
              </TableCell>
              }
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Box>
  );
};

export default CurrentSignoutsDisplay;
