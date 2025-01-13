package helper

import (
	"log"

	"gorm.io/gorm"
)
func CommitOrRollback(tx *gorm.DB) {
	if tx == nil {
		return
	}
	if r := recover(); r != nil {
		tx.Rollback()
		log.Printf("Error during commit: %v", r)
	} else {
		if err := tx.Commit().Error; err != nil {
			log.Printf("Error during commit: %v", err)
			tx.Rollback()
		}
	}
}