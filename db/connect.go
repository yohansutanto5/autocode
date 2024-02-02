package db

import "gorm.io/gorm"

type DataStore struct {
	Db gorm.DB
}
