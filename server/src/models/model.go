package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"log"
)

type Model struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error{
	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil{
		log.Fatalf("Error during Company Object creating: %v", err)
	}
	return err
}