package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Naam     string `json:"naam"`
	Checkout string `json:"checkout"`
}

func main() {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	db, err := sql.Open("mysql", "Fonteyn:P@ssword@tcp(reserveringen.mysql.database.azure.com:3306)/klanten?tls=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/nummerplaat", func(w http.ResponseWriter, r *http.Request) {
		licenseplate := r.URL.Query().Get("licenseplate")
		if licenseplate != "" {
			info := db.QueryRow("SELECT name, checkout FROM reservering WHERE kenteken=? AND checkout >=?", licenseplate, currentDate)
			var data Data
			err = info.Scan(&data.Naam, &data.Checkout)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "licenseplate or date not valid", http.StatusNotFound)
					return
				}
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
