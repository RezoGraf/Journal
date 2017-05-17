package models

import "time"

type EditAddNar struct {
	IdPatient int64 `json:"id_patient"`
	IdVrach int64 `json:"id_vrach_ortoped"`
	IdVrachTechnic int64 `json:"id_vrach_technic"`
	Numbernar string `json:"number_nar"`
	DateOpenNar time.Time `json:"date_open_nar"`
	DateStartProduction time.Time `json:"date_start_producrion"`
	DateCloseNar time.Time `json:"date_close_nar"`
	Sum string `json:"sum"`

}
