package media

import (
	"fmt"

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

func (c Content) String() string {
	return fmt.Sprintf("%d-%s-%s[%s]", c.ID, c.Title, c.Status, c.URLName)
}
