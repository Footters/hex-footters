package mysqldb

import (
	"errors"

	"github.com/Footters/hex-footters/pkg/content"
	"github.com/jinzhu/gorm"
)

type contentRepository struct {
	db *gorm.DB
}

// NewMysqlContentRepository constructor
func NewMysqlContentRepository(db *gorm.DB) content.Repository {

	return &contentRepository{
		db: db,
	}
}

func (r *contentRepository) Create(content *content.Content) error {
	r.db.Create(content)
	return nil
}
func (r *contentRepository) FindByID(id uint) (*content.Content, error) {

	cID := content.Content{}
	cID.ID = id
	cOUT := []content.Content{}
	r.db.Where(&cID).Find(&cOUT)

	if len(cOUT) != 0 {
		return &cOUT[0], nil
	}

	return nil, errors.New("Not find")
}
func (r *contentRepository) FindAll() ([]content.Content, error) {
	allContents := []content.Content{}
	r.db.Find(&allContents)
	return allContents, nil
}
