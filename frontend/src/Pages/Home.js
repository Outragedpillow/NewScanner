import React from "react";
import { Box, Container } from "@mui/material";
import {Grid} from "@mui/material";

export default function Home() {
  const [responseData, setResponseData] = React.useState({});

  const postApiData = async (data) => {
    try {    
      const response = await fetch('http://localhost:5432/api/check-scan', {
        mode: 'cors',
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({scan: data}),
      })
      
      if (response.ok) {
        const result = await response.json();
        setResponseData(result);
      } else {
        setResponseData({"Error": response.text});
      }
    } catch (error) {
       console.error(error);
    }
  }

  return (
    <Container>
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <h1>Scan Data</h1>
          <DataInput sendData={postApiData} />
          <DisplayScanResponse responseData={responseData} />
        </Grid>
        
        <Grid item xs={6}>
          <h1>Current Signouts</h1>
          {/* Place your content for the second box here */}
        </Grid>
      </Grid>
    </Container>
  )
}

const DataInput = ({sendData}) => {
  const [inputData, setInputData] = React.useState("");
  

  const handleInputChange = (event) => {
    setInputData(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    sendData(inputData);
    setInputData("");
  }


  return (
      <form onSubmit={handleSubmit}>
        <input value={inputData} autoFocus placeholder="Scan" onChange={handleInputChange} />
      </form>
  )
}

const DisplayScanResponse = ({ responseData }) => {
  const renderObject = (obj) => {
    if (obj === null) {
      return <span>null</span>;
    }

    return (
      <ul>
        <li>{responseData.Name}</li>
        {Object.entries(obj).map(([key, value]) => (
          <li key={key}>
            <strong>{key}:</strong> {renderValue(value)}
          </li>
        ))}
      </ul>
    );
  };

  const renderArray = (arr) => {
    return (
      <ul>
        {arr.map((item, index) => (
          <li key={index}>{renderValue(item)}</li>
        ))}
      </ul>
    );
  };

  const renderValue = (value) => {
    if (typeof value === 'object') {
      if (Array.isArray(value)) {
        return renderArray(value);
      } else {
        return renderObject(value);
      }
    } else {
      return String(value);
    }
  };

  return (
    <div>
      <h2>Response Data:</h2>
      {renderObject(responseData)}
    </div>
  );
};

const DisplayCurrentSignouts = () => {
  return (
    <h1>Name Here</h1>
  )
}
