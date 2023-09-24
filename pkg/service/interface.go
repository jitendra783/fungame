package service

import (
	"fungame/pkg/db"
	v1 "fungame/pkg/service/v1"

	"github.com/gin-gonic/gin"
)

type serObj struct {
	db db.DBGroup
	v1 v1.ServiceGroup
}

type ServiceLayer interface {
	v1.ServiceGroup
	Status(*gin.Context)	
}

func NewServiceGroup(db db.DBGroup) ServiceLayer {
	return &serObj{db: db}
}
