package repository

import (
	"math"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logrus.WithFields(logrus.Fields{"type": "repository"})

// Try to create gorm connection, until maxAttempt attempts are made. If maxAttempt is 0, tries infinitely.
func RetryConnectGorm(postgresConnection string, maxAttempt uint) (db *gorm.DB, err error) {
	if maxAttempt == 0 {
		maxAttempt = uint(math.MaxUint64)
	}
	attemptCnt := uint(0)

	for attemptCnt < maxAttempt {

		db, err = gorm.Open(postgres.Open(postgresConnection), &gorm.Config{})
		if err == nil {
			sqldb, err := db.DB()
			if err == nil {
				sqldb.SetMaxIdleConns(0)
			}

			break
		} else {
			attemptCnt++
			log.WithFields(logrus.Fields{
				"err":              err.Error(),
				"connectionString": postgresConnection,
			}).Warnf("failed to connect to database, retrying... %v", attemptCnt)
			time.Sleep(1 * time.Second)
		}
	}

	return db, err
}
