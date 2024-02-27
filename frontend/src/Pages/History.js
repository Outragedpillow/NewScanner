import { Box, Table, TableCell, TableRow, Typography } from "@mui/material";
import React from "react";
import HistoryDialog from "../Components/HistoryDialog";
import axios from "axios";

const HistoryPage = () => {
  const getCurrentDate = () => {
    const date = new Date();
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const year = String(date.getFullYear()).slice(-2);
    return `${month}/${day}/${year}`;
  };

  const [postResp, setPostResp] = React.useState([]);
  const [formData, setFormData] = React.useState({
    date: getCurrentDate().toString(),
    mdoc: "",
    serial: "",
  })

  React.useEffect(() => {
    const postData = async () => {
      try {
        const response = await axios.post("http://localhost:5432/api/history", formData);
        setPostResp(response.data.historyassignment);
      } catch (err) {
        console.log(err);
      }
    } 
    postData();
  }, [formData])

  const handleDialogSubmit = async (data) => {
    setFormData(data);
  } 

  return (
    <Box >
    <Typography variant="h1" sx={{ml: 50 }}>History</Typography>
    <HistoryDialog onSubmit={handleDialogSubmit} />
    <Table sx={{ml: 1, width: .99}}>
      <TableRow>
        <TableCell><Typography>Name</Typography></TableCell>
        <TableCell><Typography>Mdoc</Typography></TableCell>
        <TableCell><Typography>Display Name</Typography></TableCell>
        <TableCell><Typography>Device Serial</Typography></TableCell>
        <TableCell><Typography>Time Issued</Typography></TableCell>
        <TableCell><Typography>Time Returned</Typography></TableCell>
      </TableRow>
      {postResp !== null ? postResp.map((item, index) => {
        return (
          <TableRow key={index}>
            <TableCell>{item.name}</TableCell>
            <TableCell>{item.mdoc}</TableCell>
            <TableCell>{item.qr_tag}</TableCell>
            <TableCell>{item.serial}</TableCell>
            <TableCell>{item.time_issued}</TableCell>
            <TableCell>{item.time_returned}</TableCell>
          </TableRow>
      )}): <TableRow><TableCell>No History</TableCell></TableRow>}
    </Table>
    </Box>
  )
}

export default HistoryPage;
