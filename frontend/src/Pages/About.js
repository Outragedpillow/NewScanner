import { Box, Grid, Paper, Typography } from "@mui/material";
import React from "react";

const AboutPage = () => {
  return (
    <Grid container spacing={3}>
      <Grid item md={5}>
        <Paper elevation={10} sx={{ml: .5, mt: 1}}>
          <Box sx={{ml: 1}}>
            <Typography variant="h2">How To Use</Typography>
            <Typography variant="h5">Basic Steps</Typography>
            <Typography variant="body1"><strong>Step 1:</strong> Scan or Enter the Residents MDOC or Staff Qr Code</Typography>
            <Typography variant="body1"><strong>Step 2:</strong> Scan or Enter the Device Qr Code or Barcode.</Typography>
            <Typography variant="body1"><strong>Step 3:</strong> Scan or Enter the Break Qr Code</Typography>
          </Box>
          <Box sx={{pt: 5, ml: 1}}>
            <Typography variant="h5">Additional Information</Typography>
            <Typography variant="body1" sx={{mt: .5}}><strong>Step 2:</strong> Step 2 is repeated for each device and will continue to assign/unassign devices
              until the Break Qr code is entered or scanned.
            </Typography>
            <Typography variant="body1" sx={{mt: .5}}><strong>Cursor Focus:</strong> If the cursor is not focused on the input field above 
              the word SUBMIT in the sidebar, none of the scans will work. The input field is where all the scan data must be entered. If for 
              any reason, like clicking to change to show the history, another part of the page is clicked the cursor will no longer be focued 
              in the input field and any scanned data will not be entered. To refocus the cursor simply click the input field area right above 
              the word SUBMIT in the sidebar.
            </Typography>
            <Typography variant="body1"><strong>Manually Entering Data:</strong> All data, such as resident mdoc number, staff qr code, device serial, device qr code, 
              can be typed into the input field above the word SUBMIT in the sidebar after which you can hit enter or click the word SUBMIT to enter the data.
              BEWARE if you click the SUBMIT button you must then refocus the cursor on the input field (see Cursor Focus).
            </Typography>
            <Typography variant="body1"><strong>Staff Qr codes:</strong> A device can be assigned/unassigned to a staff member, e.g. a teacher, by 
              scanning the corresponding qr code for them. Like with the resident ID, staff qr code can also be manually entered, i.e. typed in, 
              by entering their code. Each staff person's qr code is written by their name above their qr code.
            </Typography>
            <Typography variant="body1"><strong>Scanning the Break Qr:</strong> This step is extremely important. This step tells the program
              that you are done scanning devices and that it should expect a resident and or staff qr next.
              If the Break Qr is not scanned the next resident ID/Staff Qr scanned will not register because the program will think you are still
              scanning devices.
            </Typography>
          </Box>
        </Paper>
      </Grid>
      <Grid item md={6.5}>
        <Paper elevation={0} sx={{mt: 1}}>
          <Typography variant="h2">Common Mistakes</Typography>
            <Typography variant="body1"><strong>Inactive Scanner State:</strong> After a period of inactivity, the scanner will go into a sleep
              mode. Pulling the trigger will reactivate the scanner but if something is scanned to quickly after the scanner is reactivated 
              it may not properly scan the data. Sometimes this will cause the scanner to emit a warning beep however this will not always
              happen. To prevent this from happening it is best to pull the trigger a few times, until seeing the scanner flash at least twice,
              before scanning an item. 
            </Typography>
            <Typography variant="body1"><strong>Forgetting to Break:</strong> Forgetting to scan/enter the Break statement will result 
              in a 'sql: no rows in result error'. See Scanning the Break Qr under the Additional Information section for more details. 
            </Typography>
            <Typography variant="body1"><strong>Unfocused Cursor:</strong> Unfocused cursor refers to when the focus is not on the input field
              above the word SUBMIT in the sidebar. This is commonly caused by the user clicking another part of the screen and not clicking back
              onto the input field afterwards. One way to tell if the focus is on the input field is if within the input field there is a blinking
              cursor. Signs that an unfocused cursor is the issue is if when a scan occurs nothing happens on the screen. To fix this issue simply
              click on the area directly above the word SUBMIT in the top part of the sidebar.
            </Typography>
            <Typography variant="body1"><strong>Disconnected Scanner Receiver:</strong> If the cursor is properly focused and 
              nothing is happening when scanning then it is likely that the receiver for the scanner is not properly connected. 
              If the receiver is not inserted into the computer, simply insert it into the usb port. Otherwise Try removing the 
              receiver then reinserting it into the usb port. 
            </Typography>
        </Paper>
      </Grid>
    </Grid>
  )
}

export default AboutPage;
