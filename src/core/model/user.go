package model

const (
	UserRole  = "User"
	AdminRole = "Admin"
)

type User struct {
	Role  string `gorm:"not null"`
	Token string `gorm:"not null"`
}
