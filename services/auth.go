package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/techswarn/backend/database"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserDetails struct {
	Id string
	Email string
	UserName string
	FirstName string
}

type data struct {
	Token string
	User UserDetails
}
// Signup returns JWT token for the user
func Signup(userInput models.UserRequest) (string, error) {
    // create a password using bcrypt library
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

    // if password creation is failed, return the error
	if err != nil {
		return "", err
	}
	fmt.Printf("User input in services %#v", userInput)
    // create a new user object
    // this user will be added into the database
	var user models.User = models.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		FirstName: userInput.FirstName,
		LastName: userInput.LastName,
		Access: userInput.Access,
		Phone: userInput.Phone,
		Password: string(password),
		Type: userInput.Type,
	}

    // create a user into the database
	dbuser := database.DB.Create(&user)

	fmt.Printf("Create user DB error: %s \n", dbuser.Error)

	if dbuser.Error != nil {
		return "", errors.New("Phone number already exists")
	}

	address := models.Address{
		State: userInput.State,
		City:	userInput.City,
		Street: userInput.Street,
		Pincode: userInput.Pincode,
		User_id: user.ID,
	}

	res := database.DB.Create(&address)

	fmt.Printf("DB error: %s \n", res.Error)

    // generate the JWT token
	token, err := utils.GenerateNewAccessToken(user.ID)

    // if generation is failed, return the error
	if err != nil {
		return "", err
	}

    // return the JWT token
	return token, nil
}

// Login returns JWT Token for the registered user
func Login(userInput models.UserLoginRequest) (data, error) {
    // create a variable called "user"
	var user models.User

    // find the user based on the email
	result := database.DB.First(&user, "email = ?", userInput.Email)
	fmt.Printf("%s", user.ID)
    // if the user is not found, return the error
	if result.RowsAffected == 0 {
		return data{}, errors.New("user not found")
	}

    // compare the password input with the password from the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

    // if the password is not match, return the error
	if err != nil {
		return data{}, errors.New("invalid password")
	}

    // generate the JWT token
	token, err := utils.GenerateNewAccessToken(user.ID)
   
    // if generation is failed, return the error
	if err != nil {
		return data{}, err
	}

	res := data{
		Token: token,
		User: UserDetails{
			Id: user.ID,
			Email: user.Email,
			UserName: user.UserName,
			FirstName: user.FirstName,
		},

	}

    // return the JWT token
	return res, nil;
}
