package services

import (
	"fmt"
	"github.com/techswarn/backend/models"
    "github.com/google/uuid"
    "github.com/techswarn/backend/database"
)

func CreateTag(tagRequest models.Tag_Request) (models.Tag, error) {
	var tag models.Tag = models.Tag{
		TagID: uuid.New().String(),
		Name: tagRequest.Name,
	}

	if result := database.DB.Create(&tag); result.Error != nil {
		fmt.Printf("DB write error: %s", &result.Error)
		return tag, result.Error
	}

	return tag, nil
}

func GetTag() ([]models.Tag, error) {
	var Tags []models.Tag = []models.Tag{}

	if result := database.DB.Order("created_at desc").Find(&Tags); result.Error != nil {
		fmt.Printf("DB write error: %s", &result.Error)
		return Tags, result.Error
	}

	return Tags, nil
}