package main

import (
    "net/http"
)

func aboutHandler(wr http.ResponseWriter, r *http.Request) {
    wr.Write([]byte("About Page"))
}
func indexHandler(wr http.ResponseWriter, r *http.Request) {
    wr.Write([]byte("İndex Page"))
    //wr.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/about", aboutHandler)
    http.ListenAndServe(":8080", nil)
}

/*
     http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        rw.Write([]byte("Merhaba Mars!"))
    })
*/
