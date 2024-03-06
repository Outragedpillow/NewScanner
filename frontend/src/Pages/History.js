import { Box, Table, TableCell, TableRow, Typography } from "@mui/material";
import React from "react";
import HistoryDialog from "../Components/HistoryDialog";
import axios from "axios";
import { DataGrid } from "@mui/x-data-grid";

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

  const rowsWithIds = postResp.map((item, index) => {
    return { id: index + 1, ...item };
  });


  const columns = [
    {field: "name", headerName: "Name", width: 200},
    {field: "mdoc", headerName: "MDOC", width: 200},
    {field: "qr_tag", headerName: "Display Name", width: 200},
    {field: "serial", headerName: "Device Serial", width: 200},
    {field: "time_issued", headerName: "Time Issued", width: 200},
    {field: "time_returned", headerName: "Time Returned", width: 200},
    {
    field: "customHeader",
    headerName: "Custom Header",
    width: 250,
    renderHeader: () => <HistoryDialog onSubmit={handleDialogSubmit} />,
  },
  ]

  return (
    <Box sx={{ height: "95vh", overflow: "auto", ml: 0.5 }}>
      <DataGrid sx={{height: "94vh"}} columns={columns} pagination pageSize={5} rows={rowsWithIds} initialState={{
          pagination: {
            paginationModel: { page: 0, pageSize: 5 },
          },
        }}
        pageSizeOptions={[5, 10, 15, 20, 25, 50, 100]}
 />
    </Box>

  )
}

export default HistoryPage;
