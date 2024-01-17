package services

import (
	"fmt"
	"github.com/techswarn/backend/models"
    "github.com/google/uuid"
    "time"
    "github.com/techswarn/backend/database"
)

func CreateBlog(blogRequest models.Blog_request) (models.Blog, error) {

	var blog models.Blog = models.Blog{
		ID: uuid.New().String(),
		UserID: blogRequest.UserID,
		Subject: blogRequest.Subject,
		Paragraph: blogRequest.Paragraph,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("%#v", blog)
	
	if result := database.DB.Create(&blog); result.Error != nil {
		fmt.Printf("DB write error: %s", &result.Error)
		return blog, result.Error
	}

	return blog, nil
}

func GetAllBlogs() []models.Blog {
	// create a variable to store items data
	var Blogs []models.Blog = []models.Blog{}

	// get all data from the database order by created_at
	database.DB.Order("created_at desc").Find(&Blogs)

	// return all items from the database
	return Blogs
}

func GetBlogs(keyword string, tag string ) []models.Blog {
	// create a variable to store items data
	var Blogs []models.Blog = []models.Blog{}
    fmt.Println(keyword)
	database.DB.Where("subject LIKE ?", "%"+keyword+"%").Find(&Blogs)
   // SELECT * FROM users WHERE name LIKE '%jin%';

	// return all items from the database
	return Blogs
}