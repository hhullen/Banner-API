package model

const (
	UserRole  = "User"
	AdminRole = "Admin"
)

type User struct {
	role  string `gorm:"not empty"`
	token string `gorm:"not empty"`
}
