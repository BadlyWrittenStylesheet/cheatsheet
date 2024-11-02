package db

import (
	"database/sql"
	"fmt"

	// "fmt"
	"os"

	// "log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
    Conn *sql.DB
}

func (db *Database) PopulateDB() error {
    dir, err := os.Open("backend/cheatsheets")
    if err != nil {
        return err
    }

    files, err := dir.Readdir(-1)
    if err != nil {
        return err
    }

    fmt.Println(os.Getwd())

    for _, f := range files {
        db.Conn.Exec("INSERT INTO cheatsheets (title, file_path) VALUES (?, ?)", f.Name(), "backend/cheatsheets/" + f.Name())
    }

    return nil
}

func InitDB(dataSourceName string) (*Database, error) {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec("DROP TABLE IF EXISTS cheatsheets")
    if err != nil {
        return nil, err
    }

    queryCreateTable := `CREATE TABLE IF NOT EXISTS cheatsheets (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        file_path TEXT NOT NULL
    )`

    if _, err := db.Exec(queryCreateTable); err != nil {
        return nil, err
    }

    return &Database{Conn: db}, nil
}

func (db *Database) CreateCheatsheet(title, file_path string) (error) {
    _, err := db.Conn.Exec("INSERT INTO cheatsheets (title, file_path) VALUES (?, ?)", title, file_path)
    if err != nil {
        return err
    }
    return nil
}

type Cheatsheet struct {
    ID          int     `json:"id"` 
    Title       string  `json:"title"`
    FilePath    string  `json:"file_path"`
}

func (db *Database) GetAllCheatsheets() ([]Cheatsheet, error) {
    rows, err := db.Conn.Query("SELECT id, title, file_path FROM cheatsheets")
    if err != nil {
        return nil, err
    }

    var cheatsheets []Cheatsheet

    for rows.Next() {
        var id int
        var title, filePath string

        err = rows.Scan(&id, &title, &filePath)
        if err != nil {
            return nil, err
        }

        cheatsheet := Cheatsheet{
            ID: id,
            Title: title,
            FilePath: filePath,
        }
        cheatsheets = append(cheatsheets, cheatsheet)
    }
    return cheatsheets, nil
}

func (db *Database) GetCheatsheet(id int) (string, string, error) {
    var title, file_path string
    err := db.Conn.QueryRow("SELECT title, file_path FROM cheatsheets WHERE id = ?", id).Scan(&title, &file_path)
    if err != nil {
        return "", "", err
    }
    return title, file_path, nil
}

// I guess i could implement the other crud stuff, but who cares, and who would want to write more javascript *vomits* just to handle that.

func (db *Database) Close() error {
    return db.Conn.Close()
}


