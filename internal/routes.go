package internal

import (
	"server/test-api/internal/app"
	"server/test-api/internal/users"
)

func HandleRoutes() {
	app.PingHandler()
	users.UsersHandler()
}
