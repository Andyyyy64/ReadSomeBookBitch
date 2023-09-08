package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string     `json:"username" gorm:"unique"`
	Email      string     `json:"email" gorm:"unique"`
	Password   string     `json:"password"`
	Categories []Category `json:"categories" gorm:"foreignKey:UserID"`
	Books      []Books    `json:"books" gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	Name          string     `json:"name"`
	UserID        uint       `json:"user_id"`
	Books         []Books    `json:"books" gorm:"foreignKey:CategoryID"`
	ParentID      *uint      `gorm:"column:parent_id"`
	SubCategories []Category `json:"sub_category" gorm:"foreignKey:ParentID"`
}

type Books struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
	CategoryID  int    `json:"category_id"`
	UserID      uint   `json:"user_id"`
}
