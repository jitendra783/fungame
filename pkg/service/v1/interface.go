package v1

import (
	"fungame/pkg/db"
)

type serObj struct {
	db db.DBGroup
}
type ServiceGroup interface {
}

func NewService(db db.DBGroup) ServiceGroup {
	return &serObj{db: db}
}
