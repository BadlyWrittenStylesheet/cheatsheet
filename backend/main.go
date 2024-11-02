package main

import (
	"cheatsheet/backend/db"
	"cheatsheet/backend/handlers"
	// "fmt"
	"log"
	"net/http"
)

func main() {
    database, err := db.InitDB("cheatsheet.db")
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer database.Close() // So simple

    err = database.PopulateDB()
    if err != nil {
        log.Fatalf("Failed to populate db: %v", err)
    }
    handler := &handlers.Handler{DB: database}

    http.HandleFunc("/cheatsheets", handler.GetAllCheatsheets)
    http.HandleFunc("/cheatsheets/", handler.GetCheatsheetById)

    port := ":55003"
    log.Printf("Server listening on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}


    // rows, err := database.Conn.Query("SELECT * FROM cheatsheets")
    // if err != nil {
    //     log.Fatalf("Failed to select rows lmao")
    // }
    // cols, err := rows.Columns()
    // if err != nil {
    //     log.Fatalf("Failed to fetch rows")
    // }
    // fmt.Println(cols)

    // for rows.Next() {
    //     var id int
    //     var name, filePath string

    //     err := rows.Scan(&id, &name, &filePath)
    //     if err != nil {
    //         log.Fatalf("failed to scan row %v", err)
    //     }

    //     fmt.Printf("%d, %s, %s\n", id, name, filePath)
    // }

