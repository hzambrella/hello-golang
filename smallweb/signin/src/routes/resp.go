package routes

import (
	"io"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func String(w http.ResponseWriter, code int, result string) {
	w.WriteHeader(code)
	w.Header().Set("codelog", strconv.Itoa(code))
	io.WriteString(w, result)
}
