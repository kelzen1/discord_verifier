package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var dbPtr *gorm.DB

var (
	mutex sync.Mutex
	once  sync.Once
)

func initOnce() {
	log.Println("[database] start")

	databaseUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	dbPtr = db

	if err != nil {
		log.Panicln("[database] cannot connect:", err)
		return
	}

	log.Println("[database] connected")
}

// Get singleton
func Get() *gorm.DB {

	mutex.Lock()
	defer mutex.Unlock()

	once.Do(initOnce)

	return dbPtr
}

func scan[T any](query *gorm.DB) (result []T, err error) {
	result = []T{}

	rows, err := query.Rows()

	if err != nil {
		return
	}

	for rows.Next() {
		var tmp T
		err = query.ScanRows(rows, &tmp)

		if err != nil {
			return
		}

		result = append(result, tmp)
	}

	if len(result) == 0 {
		err = sql.ErrNoRows
	}

	return
}

func ScanMany[T any](query *gorm.DB) (result []T, err error) {
	return scan[T](query)
}
