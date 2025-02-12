package app

import (
	"net/http"
	"server/test-api/internal/utils"
)

func PingHandler() {
	http.HandleFunc("/ping", Ping)
	utils.PrintRoute("/ping")
}
