package user

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"type:varchar(100);"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Password  string    `json:"-" gorm:"type:varchar(255);"`
	Points    int       `json:"points" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) AddPoints(points int) {
	u.Points += points
}
