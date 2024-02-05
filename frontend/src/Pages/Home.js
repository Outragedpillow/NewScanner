import React from "react";
import Sidebar from "../Components/Sidebar";
import CurrentSignoutsDisplay from "../Components/CurrentSignoutsDisplay";
import { Grid, Paper, Typography } from "@mui/material";
import "../App.css";

const Home = () => {
  const [getData, setGetData] = React.useState([]);
  const [postData, setPostData] = React.useState({});
  const [clearScans, setClearScans] = React.useState(false);

  return (
    <>
      <Grid container spacing={1}>
        <Grid item md={3}><Paper elevation={20} sx={{mt: 2.5, ml: 1.5}}><Sidebar setGetData={setGetData} postData={postData} setPostData={setPostData} setClearScans={setClearScans} clearScans={clearScans} /></Paper></Grid>
        <Grid item md={9}><Paper sx={{mt: 2.5, mr: 1.5}}><CurrentSignoutsDisplay currentSignouts={getData} /></Paper></Grid>
      </Grid>
    </>
  )
}

export default Home;
