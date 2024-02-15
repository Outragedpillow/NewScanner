package handlers

import (
	"NewScanner/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HandleGetHistory(w http.ResponseWriter, r *http.Request, db *structs.Database) {
    var postData structs.HistoryPostData;
    var history []structs.HistoryAssignment;

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK);
        return;
    }

    if r.Method != http.MethodPost {
        response := structs.ScanResponse{
            Success: false,
            Type:    "ERROR",
            Action:  "GET History",
            Error: structs.Error{
                Place:   "history.go HandleGetHistory Method != POST",
                Message: "Method != POST. Invalid request type.",
            },
        }

        encodeErr := json.NewEncoder(w).Encode(response);
        if encodeErr != nil {
            http.Error(w, "HandleGetHistory Failed to encode response to json.", http.StatusInternalServerError);
        }

        return;
    }

    postHistoryDataErr := json.NewDecoder(r.Body).Decode(&postData)
    if postHistoryDataErr != nil {
      response := structs.ScanResponse {
        Success: false,
        Type: "ERROR",
        Action: "POST Form Data",
        Error: structs.Error{
          Place: "HandleGetHistory postHistoryDataErr. Failed to decode json",
          Message: postHistoryDataErr.Error(),
        },
      }

      encodeErr := json.NewEncoder(w).Encode(response);
      if encodeErr != nil {
        http.Error(w, "HandleGetHistory GetPostData. Failed to encode json", http.StatusInternalServerError);
      }
    }

    var query string;

    if postData.IsNil() {
      formattedDate := time.Now().Format("01/02/06");
      query = fmt.Sprintf("select resident_mdoc, resident_name, device_type, device_serial, time_issued, time_returned from assignments where day = '%s';", formattedDate);
    } else {
      query = postData.BuildQuery();
    }

    rows, queryErr := db.Conn.Query(query);
    if queryErr != nil {
        http.Error(w, "Query Err", http.StatusInternalServerError);
        return;
    }

    for rows.Next() {
        var assignment structs.HistoryAssignment;

        scanErr := rows.Scan(&assignment.Mdoc, &assignment.Name, &assignment.Type, &assignment.Serial, &assignment.Time_issued, &assignment.Time_returned);
        if scanErr != nil {
            response := structs.ScanResponse{
                Success: false,
                Type:    "ERROR",
                Action:  "Scan Rows",
                Error: structs.Error{
                    Place:   "history.go HandleGetHistory scanErr != nil",
                    Message: scanErr.Error(),
                },
            }

            encodeErr := json.NewEncoder(w).Encode(response);
            if encodeErr != nil {
                http.Error(w, "HandleGetHistory scanErr != nil. Failed to encode json", http.StatusInternalServerError);
            }
            return;
        }
        history = append(history, assignment);
    }

    response := structs.ScanResponse{
        Success:      true,
        RefreshCurr:  false,
        Type:         "HISTORY",
        Action:       "Return History",
        History: history,
    }

    err := json.NewEncoder(w).Encode(response);
    if err != nil {
        http.Error(w, "HandleGetHistory Failed to encode response to json.", http.StatusInternalServerError);
    }
}
