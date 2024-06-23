package schemas

import "time"

type UserSchema struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Permissions uint      `json:"permissions"`
}
