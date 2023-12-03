package utils

import "gorm.io/gorm"

func CommitRollback(db *gorm.DB) {
	err := recover()
	if err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}
