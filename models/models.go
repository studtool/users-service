package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type User struct {
	Id          string          `json:"userId"`
	Username    string          `json:"username"`
	FullName    *string         `json:"fullName"`
	DateOfBirth *types.Date     `json:"dateOfBirth"`
	Location    *LocationInfo   `json:"locationInfo"`
	University  *UniversityInfo `json:"universityInfo"`
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
	Country string  `json:"country"`
	City    *string `json:"city"`
}

//easyjson:json
type UniversityInfo struct {
	Name           string  `json:"name"`
	Department     *string `json:"department"`
	Speciality     *string `json:"speciality"`
	AdmissionYear  *int    `json:"admissionYear"`
	GraduationYear *int    `json:"graduationYear"`
}
