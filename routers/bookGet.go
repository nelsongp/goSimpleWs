package routers

import (
	"encoding/json"
	"github.com/nelsongp/goSimpleWs/books"
	"net/http"
)

func BookGet(w http.ResponseWriter, r *http.Request){
	bookID := r.URL.Query().Get("id")

	b := books.GetBook(bookID)

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}
