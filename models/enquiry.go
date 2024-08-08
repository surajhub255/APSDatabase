package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enquiry struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string         `json:"name"`
	Email		string		   `json:"email"`
	Number      int            `json:"number"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
}

func (c *Enquiry) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
