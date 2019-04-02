package mysqldb

import (
	"errors"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/jinzhu/gorm"
)

type contentRepository struct {
	db *gorm.DB
}

// NewMysqlContentRepository constructor
func NewMysqlContentRepository(db *gorm.DB) media.ContentRepository {

	return &contentRepository{
		db: db,
	}
}

func (r *contentRepository) Create(content *media.Content) error {
	r.db.Create(content)
	return nil
}
func (r *contentRepository) FindByID(id uint) (*media.Content, error) {

	cID := media.Content{}
	cID.ID = id
	cOUT := []media.Content{}
	r.db.Where(&cID).Find(&cOUT)

	if len(cOUT) != 0 {
		return &cOUT[0], nil
	}

	return nil, errors.New("Not find")
}
func (r *contentRepository) FindAll() ([]media.Content, error) {
	allContents := []media.Content{}
	r.db.Find(&allContents)
	return allContents, nil
}

func (r *contentRepository) Update(content *media.Content) error {
	r.db.Save(&content)
	return nil
}
