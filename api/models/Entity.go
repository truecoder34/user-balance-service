package models

import (
	"time"

	//"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

/*
	THE HOOK.
	Method will be called BEFORE each create call in ORM.
 	And will generate UUID for ID
*/
func (base *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()

	// if !base.IsValid() {
	// 	return errors.New("rollback invalid user")
	//   }
	return
}
