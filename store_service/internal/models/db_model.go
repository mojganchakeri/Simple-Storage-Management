package models

import "gorm.io/gorm"

type SqlClient struct {
	DB *gorm.DB
}
