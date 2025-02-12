package users

import (
	"encoding/json"
	"net/http"
	"server/test-api/internal/models"
)

func FindUserById(
	w http.ResponseWriter,
	r *http.Request) {
	idString := r.PathValue("id")
	user, err := GetUserById(idString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": http.StatusNotFound, "message": "User not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func FindUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	list, err := ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)

}

func seedUsersExample(
	w http.ResponseWriter,
	r *http.Request,
) {
	list, err := ListUsers()
	if err != nil || len(list) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"status": http.StatusInternalServerError, "message": "Users already seeded"}
		json.NewEncoder(w).Encode(response)
		return
	}

	users := []models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
		},
	}
	for _, user := range users {
		CreateUser(user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": http.StatusOK, "message": "Users seeded"})

}
