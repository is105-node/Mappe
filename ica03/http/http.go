package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    for key, value := range r.Form {
    	fmt.Fprintf(w, "{%s:%s}", key, value)
    }
}