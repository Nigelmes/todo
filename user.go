package todo

type User struct {
	Id       int    `json:"-" gorm:"column:id"`
	Name     string `json:"name" binding:"required" gorm:"column:name"`
	Username string `json:"username" binding:"required" gorm:"column:username"`
	Password string `json:"password" binding:"required" gorm:"column:password_hash"`
}
