package main

import (
	"NewScanner/handlers"
	"NewScanner/structs"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

const (
	PORT = "5432"
  STORAGE_DB = "Storage.db"
)

// Still need to handle if no Storage exists
// Should have migrate and backup Storage.db each day or something
// Still need error if device from day past is still assigned next day

func main() {
  var database structs.Database;
  
  openErr := database.Open(STORAGE_DB);
  if openErr != nil {
    panic(openErr);
  }

  mux := http.NewServeMux();

  handlers.HandleRoutes(mux, &database);

	serveErr := http.ListenAndServe(":"+PORT, mux);
	if serveErr != nil {
		panic(serveErr)
	}
}
