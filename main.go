package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

// Fatal error handler
func checkErr(err error) {
    if err != nil {
        log.Panicln(err)
    }
}

func rbcParseHandler(w http.ResponseWriter, r *http.Request) {
    if strings.Contains(r.URL.Path, "/rbc-parse") {
        response := Response{}
        data, err := rbcParse()
        if err != nil {
            response.Status = "unsuccess"
            response.Error = err.Error()
            w.Header().Set("Content-Type", "application/json")
            js, _ := json.Marshal(response)
            w.Write(js)
            return
        }
        response.Status = "success"
        response.Data = data
        response.Error = ""
        w.Header().Set("Content-Type", "application/json")
        js, _ := json.Marshal(response)
        w.Write(js)
    }
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Root")
}

func status(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    args := os.Args[1:]
    http.HandleFunc("/", root)
    http.HandleFunc("/status", status)
    http.HandleFunc("/rbc-parse", rbcParseHandler)
    log.Fatal(http.ListenAndServe(args[0], nil))
}
