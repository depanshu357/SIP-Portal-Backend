package models

import uuid "github.com/google/uuid"

type Student struct {
	ID                     uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name                   string    `gorm:"not null"`        // name
	RollNumber             string    `gorm:"unique;not null"` // rollNumber
	Email                  string    `gorm:"unique;not null"` // email
	Department             string    `gorm:"not null"`        // department
	SecondaryDepartment    string    // secondaryDepartment
	Specialisation         string    // specialisation
	Gender                 string    `gorm:"default:''"` // gender
	DOB                    string    // dob (you can also use `time.Time` if you want date handling)
	AlternateContactNumber string    // alternateContactNumber
	CurrentCPI             string    // currentCPI
	TenthBoard             string    // tenthBoard
	TenthMarks             string    // tenthMarks
	TenthBoardYear         string    // tenthBoardYear
	EntranceExam           string    // entranceExam
	Category               string    // category
	CurrentAddress         string    // currentAddress
	Disability             string    // disability
	ExpectedGraduationYear string    // expectedGraduationYear
	Program                string    // program
	SecondaryProgram       string    // secondaryProgram
	Preference             string    // preference
	PersonalEmail          string    // personalEmail
	ContactNumber          string    // contactNumber
	WhatsappNumber         string    // whatsappNumber
	TwelfthBoardYear       string    // twelfthBoardYear
	TwelfthBoard           string    // twelfthBoard
	TwelfthMarks           string    // twelfthMarks
	EntranceExamRank       string    // entranceExamRank
	CategoryRank           string    // categoryRank
	PermanentAddress       string    // permanentAddress
	FriendsName            string    // friendsName
	FriendsContactDetails  string    // friendsContactDetails
	IsVerified             bool      `gorm:"default:false"` // IsVerified
	IsProfileVerified      bool      `gorm:"default:false"` // IsProfileVerified
}
