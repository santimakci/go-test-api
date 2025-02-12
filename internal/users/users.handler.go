package users

import (
	"net/http"
	"server/test-api/internal/utils"
)

func UsersHandler() {
	http.HandleFunc("/users", FindUsers)
	utils.PrintRoute("/users")
	http.HandleFunc("/users/{id}", FindUserById)
	utils.PrintRoute("/users/{id}")
	http.HandleFunc("/seed-users", seedUsersExample)
	utils.PrintRoute("/seed-users")

}
