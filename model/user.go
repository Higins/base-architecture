package appModel

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
	Name      string
	LastLogin sql.NullTime
}
