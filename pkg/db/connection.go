package db

import (
	"fmt"
	"log"

	"github.com/stebin13/order-srv/pkg/config"
	"github.com/stebin13/order-srv/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func InitDb(c *config.Config) Handler {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Db_Host, c.Db_User, c.Db_Password, c.Db_Name, c.Db_Port)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	db.AutoMigrate(&models.Order{})

	return Handler{DB: db}
}
