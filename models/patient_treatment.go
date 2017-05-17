package models

import (
	"database/sql"
	//"time"
)


type PatientTreatment struct {
	Id         sql.NullString `json:"Id"`
	Fam        sql.NullString `json: "Fam"`
	Name       sql.NullString `json: "Name"`
	Lastname   sql.NullString `json: "Lastname"`
	Datebirth  sql.NullString `json: "Datebirth"`
	DateRecord sql.NullString `json: "DateRecord"`
	Lgotcat    sql.NullString `json: "Lgotcat"`
}

type PatientTreatmentSrting struct {
	Id         string `json:"Id"`
	Fam        string `json: "Fam"`
	Name       string `json: "Name"`
	Lastname   string `json: "Lastname"`
	FullName   string `json"FullName"`
	Datebirth  string `json: "Datebirth"`
	DateRecord string `json: "DateRecord"`
	Lgotcat    string `json: "Lgotcat"`
	DateInvitation string `json:"DateInvintation`
}

