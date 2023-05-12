package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type Data struct {
    Naam string `json:"naam"`
}

func main() {
    db, err := sql.Open("mysql", "Fonteyn:P@ssword@tcp(reserveringen.mysql.database.azure.com:3306)/klanten?tls=true")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/nummerplaat", func(w http.ResponseWriter, r *http.Request) {
        licenseplate := r.URL.Query().Get("licenseplate")
        if licenseplate != "" {
            row := db.QueryRow("SELECT klant_naam FROM klant WHERE nummerplaat=?", licenseplate)
            var data Data
            err = row.Scan(&data.Naam)
            if err != nil {
                if err == sql.ErrNoRows {
                    http.Error(w, "License not found", http.StatusNotFound)
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
