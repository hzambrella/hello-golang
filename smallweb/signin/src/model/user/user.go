package user

import "time"

type User struct {
	UserId     int       `json:"user_id"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
}
