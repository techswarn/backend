package services

import(
	"fmt"
	"github.com/techswarn/backend/models"
    _ "github.com/google/uuid"
    _ "time"
    "errors"
    "github.com/techswarn/backend/database"
)

type UserDetail struct {
	Id string
	Email string
	UserName string
	FirstName string
}

func GetUserByID(id string) (UserDetail, error) {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)

// if the item data is not found, return an error
	if result.RowsAffected == 0 {
		return UserDetail{}, errors.New("User not found")
	}

// return the item data from the database
	res := UserDetail{
		Id: user.ID,
		Email: user.Email,
		UserName: user.UserName,
		FirstName: user.FirstName,
	}
	return res, nil
}

func GetAllSellersByType(cat string) ([]models.User, error) {

	 var Users []models.User = []models.User{}

	 // get all data from the database order by created_at
	 result := database.DB.Where("type = ?", "seller").Find(&Users)


	 if result.Error != nil {
		fmt.Printf("DB read error: %s", &result.Error)
		return nil, result.Error
	}

	 // return all items from the database
	 return Users, nil
}