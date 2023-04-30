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
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	itemId int
}
