package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type User struct {
	Id          string          `json:"userId" structs:"-"`
	Username    string          `json:"username" structs:"username"`
	FullName    *string         `json:"fullName" structs:"fullName"`
	DateOfBirth *types.Date     `json:"dateOfBirth" structs:"dateOfBirth"`
	Location    *LocationInfo   `json:"locationInfo" structs:"location"`
	University  *UniversityInfo `json:"universityInfo" structs:"university"`
}

//easyjson:json
type UserMap map[string]interface{}

//easyjson:json
type UserInfo struct {
	Id       string `json:"userId"`
	Username string `json:"username"`
}

//easyjson:json
type LocationInfo struct {
	Country string  `json:"country" structs:"country"`
	City    *string `json:"city" structs:"city"`
}

//easyjson:json
type UniversityInfo struct {
	Name           string  `json:"name" structs:"name"`
	Department     *string `json:"department" structs:"department"`
	Speciality     *string `json:"speciality" structs:"speciality"`
	AdmissionYear  *int    `json:"admissionYear" structs:"admissionYear"`
	GraduationYear *int    `json:"graduationYear" structs:"graduationYear"`
}
