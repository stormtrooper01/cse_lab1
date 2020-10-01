package main

import (
    "net/http" //to create the server
    "time" //getting current time
    "encoding/json"
)

func main () { //this defines the main function
    http.ListenAndServe(":8795", nil)//start the server at port 8795 and pass the router
    http.HandleFunc("/", getTime)if we have a request in the “/” route, you invoke the getTime()
}

func getTime (w http.ResponseWriter, r *http.Request) {

    res := indexRes{time.Now().Format(time.RFC3339)}
    jsonRes, err := json.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write(jsonRes)
}

type indexRes struct {
    Time string `json:"time"`
}
