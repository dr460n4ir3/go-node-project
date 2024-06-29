package main

import (
        "fmt"
        "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, Dr460n!")
}

// serve an html file
func htmlHandler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            name := r.FormValue("name")
            email := r.FormValue("email")
            message := r.FormValue("message")
    
            // Process the form data
            fmt.Fprintf(w, "Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)
        } else {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        }
    }

func main() {
        http.HandleFunc("/", helloHandler)
        http.HandleFunc("/html", htmlHandler)
        http.HandleFunc("/submit", submitHandler)
        http.ListenAndServe(":4444", nil)
}
