package books

import (
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/database"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/models"
)

func AddBook(user models.User, bookDetails models.Books, categoryID int) error {
	var db = database.ConnectDB()
	bookDetails.UserID = user.ID
	if categoryID == 1 { // categoryID 1 means that user in root(none category directory)
		bookDetails.CategoryID = categoryID
		if err := db.Create(&bookDetails).Error; err != nil {
			return err
		}
		return nil
	} else {
		bookDetails.CategoryID = categoryID
		if err := db.Create(&bookDetails).Error; err != nil {
			return err
		}
		return nil
	}
}

func AddCategory(user models.User, categoryDetails models.Category) error {
	var db = database.ConnectDB()
	categoryDetails.UserID = user.ID
	if err := db.Create(&categoryDetails).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUserInfo(user models.User) (*models.User, *models.Books, *models.Category, error) {
	var db = database.ConnectDB()
	var books models.Books
	var category models.Category

	if err := db.Where("userID = ?", user.ID).First(&books).Error; err != nil {
		return nil, nil, nil, err
	}
	if err := db.Where("userID = ?", user.ID).First(&category).Error; err != nil {
		return nil, nil, nil, err
	}

	return &user, &books, &category, nil

}
