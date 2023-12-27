package initizalizers

import "github.com/bugsssssss/auth-gin/models"

func SyncDatabase() {
	DB.AutoMigrate(models.User{})
}
