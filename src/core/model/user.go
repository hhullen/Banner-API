package model

const (
	UserRole  = "User"
	AdminRole = "Admin"
)

type User struct {
	ID    int32  `gorm:"primary_key"`
	Role  string `gorm:"not null;check:role <> ''"`
	Token string `gorm:"not null;check:token <> ''"`
}
