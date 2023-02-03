package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductRef int `json:"productRefId"`
	Product Product `gorm:"foreignKey:ProductRef"`
	UserRef int `json:"userRefId"`
	User User `gorm:"foreignKey:UserRef"`	
}