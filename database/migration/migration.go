package migration

import (
	"sip/database"
	"sip/models"
)

func Up() {
	// migrate if the database does not exist

	if !database.DB.Migrator().HasTable(&models.User{}) {
		database.DB.Migrator().CreateTable(&models.User{})
	}

	if !database.DB.Migrator().HasTable(&models.Student{}) {
		database.DB.Migrator().CreateTable(&models.Student{})
	}

	if !database.DB.Migrator().HasTable(&models.Recruiter{}) {
		database.DB.Migrator().CreateTable(&models.Recruiter{})
	}

	if !database.DB.Migrator().HasTable(&models.Admin{}) {
		database.DB.Migrator().CreateTable(&models.Admin{})
	}

	if !database.DB.Migrator().HasTable(&models.Otp{}) {
		database.DB.Migrator().CreateTable(&models.Otp{})
	}

}

func Down() {
	database.DB.Migrator().DropTable((&models.User{}))
	database.DB.Migrator().DropTable((&models.Student{}))
	database.DB.Migrator().DropTable((&models.Recruiter{}))
	database.DB.Migrator().DropTable((&models.Admin{}))
	database.DB.Migrator().DropTable((&models.Otp{}))
}
