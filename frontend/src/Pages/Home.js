import React from "react";
import Sidebar from "../Components/Sidebar";
import { Box, Paper } from "@mui/material";
import "../Styles/App.css";
import NavBar from "../Components/NavBar";
import { Outlet } from "react-router-dom";

export const GetDataContext = React.createContext(undefined);
export const PostDataContext = React.createContext(undefined);

const Home = () => {
 const SlideNav = {
  marginTop: 2.5,
  width: '.04',
  color: 'white',
  borderRadius: 1.25,
  backgroundColor: '#388e3c',
};
  const [getData, setGetData] = React.useState([]);
  const [postData, setPostData] = React.useState({});
  const [clearScans, setClearScans] = React.useState(false);

  return (
    <PostDataContext.Provider value={postData}>
    <GetDataContext.Provider value={getData}>
        <Box sx={{display: 'flex', flexDirection: 'row', flexWrap: 'nowrap'}}>
          <Paper variant="outlined" elevation={20} sx={{ mt: 2.5, ml: 1.5, pr: 0, mr: .5, borderColor: "#388e3c", borderRadius: 1.25 }}>
            <Sidebar
              setGetData={setGetData}
              setPostData={setPostData}
              setClearScans={setClearScans}
              clearScans={clearScans}
            />
          </Paper>
          <Paper variant="outlined" elevation={20} sx={{ mt: 2.5, mr: 1, pr: 0, width: "200vh", borderWidth: 2.5, borderColor: "#388e3c", borderRadius: 1.25}}>
          <Outlet />
          </Paper>
          <Paper elevation={20} sx={SlideNav}>
            <NavBar />
          </Paper>
    </Box>
    </GetDataContext.Provider>
    </PostDataContext.Provider >
  );
};

export default Home;
