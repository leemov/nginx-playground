package main

import (
	"fmt"
	"net/http"
)

func sampleBlockedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "AkamaiGHost")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"status":"blocked"}`))
}

func sampleOkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"OK"}`))
}

func main() {
	http.HandleFunc("/block", sampleBlockedHandler)
	http.HandleFunc("/ok", sampleOkHandler)
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		fmt.Println(err)
	}
}
