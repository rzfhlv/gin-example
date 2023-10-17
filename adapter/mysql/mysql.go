package mysql

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	mySqlDB    *sqlx.DB
	once       sync.Once
	mySqlError error
)

type MySQL struct {
	db *sqlx.DB
}

func New() (*MySQL, error) {
	once.Do(func() {
		var err error
		mySqlDB, err = sqlx.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
		if err != nil {
			mySqlError = err
		}

		err = mySqlDB.Ping()
		if err != nil {
			mySqlError = err
		}
	})

	if mySqlError != nil {
		return nil, mySqlError
	}

	return &MySQL{
		db: mySqlDB,
	}, nil
}

func (m *MySQL) GetDB() *sqlx.DB {
	return m.db
}
