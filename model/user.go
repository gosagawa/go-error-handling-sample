package model

// User ユーザ
type User struct {
	ID       int
	Name     string
	IsActive bool
}

// IsValid 有効なユーザか？
func (u *User) IsValid() bool {
	return u.IsActive
}
