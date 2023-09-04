type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Categories []Category `json:"categories" gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
	UserID uint `json:"user_id"`
	Books []Book `json:"books" gorm:"foreignKey:CategoryID"`
}

type Books struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
	Rate int `json:"rate"`
	CategoryID uint `json:"category_id"`
}