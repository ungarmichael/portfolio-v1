package routes

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/contributions", getContributions)
	// http.HandleFunc("/del", deleteContribution)
	http.HandleFunc("/json", jsonHandler)
}
