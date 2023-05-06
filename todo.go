package todo

type TodoList struct {
	Id          int    `json:"-"  gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

type UserList struct {
	Id     int `gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
	ListId int `gorm:"column:list_id"`
}

type TodoItem struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Done        bool   `json:"done" gorm:"column:done"`
}

type ListsItem struct {
	Id     int `gorm:"column:id"`
	ListId int `gorm:"column:list_id"`
	ItemId int `gorm:"column:item_id"`
}
