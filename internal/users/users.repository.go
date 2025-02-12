package users

import (
	"server/test-api/internal/database"
	"server/test-api/internal/models"
)

func ListUsers() ([]models.User, error) {
	db := database.GetDb()
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil

}

func GetUserById(id string) (models.User, error) {
	db := database.GetDb()
	var user models.User
	result := db.First(&user, "id = ?", id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	db := database.GetDb()
	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
