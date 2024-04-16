package main

import (
  "go-auth/Controller"

  "github.com/gorilla/mux"
  "net/http"

  )
func main()  {

    r := mux.NewRouter()
    r.HandleFunc("/google_callback", Controller.GoogleCallback)
    r.HandleFunc("/", Controller.GoogleLoginController)

    http.ListenAndServe(":8080", r)
}
