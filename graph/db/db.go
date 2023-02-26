package db

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SecurityDatabase struct {
	Database *gorm.DB `json:"database"`
}

func (source *SecurityDatabase) Open() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	// Migrate the schema
	err = db.AutoMigrate(&DatabaseSecurityProject{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseScanner{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseScannerAnalysis{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseProjectAssets{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseProjectParameters{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseProjectCredentials{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseUserRole{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseUser{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&DatabaseSecurityIssue{})
	if err != nil {
		return err
	}
	source.Database = db
	return nil
}
