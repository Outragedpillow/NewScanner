import { Grid } from "@mui/material";
import React from "react";

const Resident = (props) => {
  return (
    <Grid container>
      <Grid item>props.name</Grid>
      <Grid item>props.mdoc</Grid>
    </Grid>
  )
}

export default Resident;
