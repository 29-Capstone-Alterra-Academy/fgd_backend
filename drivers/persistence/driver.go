package persistence_driver

import (
	"fgd/helper/waitloop"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersistenceConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *PersistenceConfig) InitPersistenceDB() *gorm.DB {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=verify-ca sslrootcert=ca.pem TimeZone=Asia/Jakarta",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	dialect := postgres.Open(connString)

	// Wait for one minute before errors out
	db, err := waitloop.WaitPersistence(dialect, 1*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
