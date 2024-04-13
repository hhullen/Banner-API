package model

const (
	UserRole  = "User"
	AdminRole = "Admin"
)

type User struct {
	Role  string `gorm:"not empty"`
	Token string `gorm:"not empty"`
}
