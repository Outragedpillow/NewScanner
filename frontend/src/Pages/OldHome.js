import React from "react";
import { Container } from "@mui/material";
import { Grid } from "@mui/material";
import MyComponent from "./OldHomePage";

let PostResponseData = {};
let GetResponseData = {};

export default function Home() {
  const [responseData, setResponseData] = React.useState({});
  const [currentSignouts, setCurrentSignouts] = React.useState([]);

  const postApiData = async (data) => {
    try {
      const response = await fetch("http://localhost:5432/api/check-scan", {
        mode: "cors",
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ scan: data }),
      });

      const result = await response.json();

      if (response.ok) {
        if (result.refreshcurr) {
          try {
            const getResponse = await fetch(
              "http://localhost:5432/api/currentsignouts",
              {
                mode: "cors",
                method: "GET",
                headers: {
                  "Content-Type": "application/json",
                },
              }
            );

            const getResult = await getResponse.json();
            setResponseData(result);
            setCurrentSignouts(getResult.object)
            console.log(GetResponseData)

          } catch (error) {
            console.log(error);
          } finally {
            setResponseData(result);
          }
        } else {
          setResponseData(result);
        }
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Container>
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <h1>Scan Data</h1>
          <DataInput sendData={postApiData} />
          <h2>Success: {JSON.stringify(responseData.success)}</h2>
          <h2>Action: {responseData.action}</h2>
          <h2>Error: {responseData.error.message}</h2>
        </Grid>

        <Grid item xs={6}>
          <h1>Current Signouts</h1>
          <MyComponent current={currentSignouts}/>
        </Grid>
      </Grid>
    </Container>
  );
}

const DataInput = ({ sendData }) => {

  const handleInputChange = (event) => {
    setInputData(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    sendData(inputData);
    setInputData("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={inputData}
        autoFocus
        placeholder="Scan"
        onChange={handleInputChange}
      />
    </form>
  );
};
