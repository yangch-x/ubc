package svc

import (
	"UBC/api/internal/config"
	"UBC/api/library/database"
	"UBC/models"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := database.NewMySQL(&c.ORM)
	err := models.InitModel(db)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
	}
}
