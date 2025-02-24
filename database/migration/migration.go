package migration

import (
	"sip/database"
	"sip/models"
	"sip/utils"
)

func Up() {
	// migrate if the database does not exist

	if err := database.DB.Migrator().AutoMigrate(&models.User{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating User table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Student{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Student table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Recruiter{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Recruiter table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Admin{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Admin table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Otp{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Otp table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Notice{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Notice table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Event{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Event table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.File{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating File table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.JobDescription{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating JobDescription table:", err)
	}

	if err := database.DB.Migrator().AutoMigrate(&models.Applicant{}); err != nil {
		utils.Logger.Sugar().Errorf("Error migrating Applicant table:", err)
	}

}

func Down() {
	database.DB.Migrator().DropTable((&models.User{}))
	database.DB.Migrator().DropTable((&models.Student{}))
	database.DB.Migrator().DropTable((&models.Recruiter{}))
	database.DB.Migrator().DropTable((&models.Admin{}))
	database.DB.Migrator().DropTable((&models.Otp{}))
	database.DB.Migrator().DropTable((&models.Notice{}))
	database.DB.Migrator().DropTable((&models.Event{}))
	database.DB.Migrator().DropTable((&models.File{}))
	database.DB.Migrator().DropTable((&models.JobDescription{}))
	database.DB.Migrator().DropTable((&models.Applicant{}))
}
