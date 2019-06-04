package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}
 
func main() {
	http.HandleFunc("/", handler)
	var port = os.Getenv("PORT")
	if port == "" {
		// no env variable for port
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
