import { Box, Typography } from "@mui/material";
import React from "react";
import { DataGrid } from "@mui/x-data-grid";
import { GetDataContext } from "../Pages/Home";

const renderDevicesCell = (params) => (
  <div style={{ display: "flex" }}>
    {params.value.map((device, index) => (
      <div key={index} style={{ marginRight: "10px" }}>{device.qr_tag}</div>
    ))}
  </div>
);

const CurrentSignoutsDisplay = () => {
  const currentSignouts = React.useContext(GetDataContext);

  const rowsWithIds = currentSignouts.map((item, index) => {
    return { id: index + 1, ...item };
  });

  if (currentSignouts === undefined) {
    return (
      <Box sx={{ height: "95vh", overflow: "auto" }}>
        <Typography variant="h3">No Devices Assigned</Typography>
      </Box>
    );
  }

  const columns = [
    { field: "name", headerName: "Name", width: 200 },
    { field: "mdoc", headerName: "Mdoc", width: 150 },
    {
      field: "devices",
      headerName: "Devices",
      width: 400,
      renderCell: renderDevicesCell,
    },
  ];

  return (
    <Box sx={{ height: "95vh", overflow: "auto", ml: 0.5 }}>
      <DataGrid columns={columns} pagination pageSize={5} rows={rowsWithIds} initialState={{
          pagination: {
            paginationModel: { page: 0, pageSize: 5 },
          },
        }}
        pageSizeOptions={[5, 10, 15, 20, 25]}
 />
    </Box>
  );
};

export default CurrentSignoutsDisplay;
