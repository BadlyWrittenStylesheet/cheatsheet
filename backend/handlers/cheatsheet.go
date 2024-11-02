package handlers

import (
	"cheatsheet/backend/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Handler struct {
    DB *db.Database
}

func mdToHTML(md []byte) []byte {
    extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
    p := parser.NewWithExtensions(extensions)
    doc := p.Parse(md)

    htmlFlags := html.CommonFlags | html.HrefTargetBlank
    opts := html.RendererOptions{Flags: htmlFlags}
    renderer := html.NewRenderer(opts)
    return markdown.Render(doc, renderer)
}

func AllowCors(w http.ResponseWriter) http.ResponseWriter {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    return w
}

func (h *Handler) GetCheatsheetById(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    log.Println("PATH:", parts)
    var cheatsheetID int;
    if len(parts) == 3 && parts[1] == "cheatsheets" && parts[2] != "" {
        var err error;
        cheatsheetID, err = strconv.Atoi(parts[2])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
        }
    }

    _, filePath, err := h.DB.GetCheatsheet(cheatsheetID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Cheatsheet with id=%d not found", cheatsheetID), http.StatusNotFound)
        return
    }

    content, err := os.ReadFile(filePath)
    if err != nil {
        http.Error(w, "Failed to read file contents", http.StatusInternalServerError)
        return
    }

    html := mdToHTML(content)

    w = AllowCors(w)

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusNoContent)
        return
    }

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(html)
}

func (h *Handler) GetAllCheatsheets(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    cheatsheets, err := h.DB.GetAllCheatsheets()
    if err != nil {
        log.Fatalf("Failed to get all cheatsheets: %v", err)
    }

    w = AllowCors(w)

    log.Println(cheatsheets)

    json.NewEncoder(w).Encode(cheatsheets)
}

