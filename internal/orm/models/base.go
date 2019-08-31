package models

import (
	"time"
	
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// BaseModel defines common columns that all db struct
// should hold, usually
// db struct base on this have no soft delete
type BaseModel struct {
	// ID should use uuid_generate_v4() for primary key
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgres)SQL
	UpdatedAt time.Time `gorm:"index"`
}

// BaseModelSoftdelete defined the common columns that
// all db struct should hold, usually
// This struct also defines the fields for GORM triggers
// to detect the entity should soft delete
type BaseModelSoftdelete struct {
	BaseModel
	DeletedAt time.Time `sql:"index"`
}