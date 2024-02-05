import React from "react";
import axios from "axios";
import { Box, Button, FormControl, Input, Table, TableBody, TableCell, TableHead, TableRow } from "@mui/material";
import { makeStyles } from '@mui/styles';

const useStyles = makeStyles({
  redBackground: {
    background: 'red',
  },
  greenBackground: {
    background: 'green'
  },
  flashBackground: {
    background: 'red',
  },
});

const Sidebar = ({ setGetData, setPostData, postData, setClearScans, clearScans }) => {
  const classes = useStyles();

  const [inputData, setInputData] = React.useState("");
  const [scanArray, setScanArray] = React.useState([]);
  const [flashColor, setFlashColor] = React.useState('');

  React.useEffect(() => {
    const getData = async () => {
      try {
        const getResponse = await axios.get("http://localhost:5432/api/currentsignouts");
        setGetData(getResponse.data.object);
      } catch (error) {
        console.log(error);
      }
    }
    getData();
  }, [setGetData]);

  React.useEffect(() => {
    if (postData.success === false) {
      setFlashColor('red');
      setTimeout(() => {
        setFlashColor('');
      }, 1000); // Reset flash color after 1 second
    } else {
      setFlashColor('');
    }
  }, [postData.success]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setInputData("");
    try {
      const postResponse = await axios.post('http://localhost:5432/api/check-scan', {scan: inputData});
      if (postResponse.data.refreshcurr === true) {
        const getResponse = await axios.get('http://localhost:5432/api/currentsignouts');
        setGetData(getResponse.data.object);
        setScanArray((prev) => [...prev, postResponse.data]);
        setClearScans(true);
      } else {
        setPostData(postResponse.data);
        if (clearScans) {
          setScanArray([]);
          setScanArray((prev) => [...prev, postResponse.data]);
          setClearScans(false);
        } else {
          setScanArray((prev) => [...prev, postResponse.data]);
        }
      }
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <Box sx={{ height: "95vh", overflow: "auto" }} className={flashColor === 'red' ? classes.flashBackground : ''}>
      <FormControl component="form" onSubmit={handleSubmit}>
        <Input autoFocus value={inputData} onChange={(e) => setInputData(e.target.value)}/>
        <Button type="submit">Submit</Button>
      </FormControl>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Success</TableCell>
            <TableCell>Found</TableCell>
            <TableCell>Action</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {scanArray.map((item) => (
            <TableRow className={item.success ? '' : classes.redBackground}>
              <TableCell>{item.success.toString().toUpperCase()}</TableCell>
              {item.object !== null ? 
                <TableCell>
                  {item.type === "Resident" ? item.object.name : item.type === "DEVICE" ? item.object.type.toUpperCase() + " " + item.object.tag_number : "" }
                </TableCell>
              : <TableCell>Null</TableCell>}
              <TableCell>{item.success ? item.action : item.error.message}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Box>
  )
}

export default Sidebar;
