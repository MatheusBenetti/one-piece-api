package user

type User struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	username string
	password string
	email    string
}
