package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DomainName struct {
	gorm.Model

	Name     string
	Address  string
	LastRead time.Time
}
