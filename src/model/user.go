package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Age       uint      `json:"age" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *User) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}

func (p *User) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *User) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *User) Updates() (tx *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *User) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *User) DeleteById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&p)
}
