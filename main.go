package main

import (
	"net/http"
)


func main() {
	runAPI()

}


func runAPI(){
	logInfo(`API running on :`+API_PORT)

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/add/", handle404)

	http.ListenAndServe(`:` + API_PORT, nil)

	logInfo(`Stopping.....`)
}





