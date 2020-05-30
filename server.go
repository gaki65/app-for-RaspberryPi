package main

import (
  "net/http"
)

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
  http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view/"))))
  http.HandleFunc("/home", home)
  server.ListenAndServe()
}
