package model

type Formula struct {
    ID int `gorm:"AUTO_INCREMENT;primary_key"`
	UserId int `gorm:"not null;"`
    Formula string `gorm:"not null;"`
    Result string `gorm:"not null;"`
}