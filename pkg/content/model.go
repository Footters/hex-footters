package content

import (
	"github.com/jinzhu/gorm"
)

// Content Model
type Content struct {
	gorm.Model
	URLName     string `json:"urlName"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Free        int    `json:"free"`
	Visible     int    `json:"visible"`
}

// Repository Content interface
type Repository interface {
	Create(content *Content) error
	FindByID(id uint) (*Content, error)
	FindAll() ([]Content, error)
	Update(content *Content) error
}
