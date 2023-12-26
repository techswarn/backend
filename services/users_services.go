package services

import(
	_ "fmt"
	"github.com/techswarn/backend/models"
    _ "github.com/google/uuid"
    _ "time"
    "errors"
    "github.com/techswarn/backend/database"
)

func GetUserByID(id string) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)

// if the item data is not found, return an error
	if result.RowsAffected == 0 {
	return models.User{}, errors.New("User not found")
	}

// return the item data from the database
	return user, nil
}