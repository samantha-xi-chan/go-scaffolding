package model

import "time"

type User struct {
	Id        string `gorm:"primaryKey"`
	Username  string
	Email     string `json:"email" gorm:"type:varchar(255);unique_index"`
	CreatedAt time.Time
}
