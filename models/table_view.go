package models

import "time"

type Table_view struct {
	Id_d		string
	Id             int64
	Fam            string
	Name           string
	Lastname       string
	Datebirth      string
	DateRecord     string
	Phone          string
	Homeadres      string
	Numberud       string
	Lgotcat        string
	Fiovrach       string
	DateInvitation string
	DateNar        string `json:"DateNar"`
	Numbernar      string
	Fioreg         string
	Comment        string
	NumberNar      string	`json:"NumberNar"`
	DateOpenNar	string
	DateCloseNar	string
	FlagPatientComplite bool
	FlagPatientRefuse bool
	DateLastEdit	time.Time
}
