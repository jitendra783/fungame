package db

import (
	"gorm.io/gorm"
)

type dbObj struct {
	db *gorm.DB
}

func NewDBservice(psql, msql *gorm.DB) DBGroup {
	return &dbObj{db: psql}
}

type DBGroup interface {
}
