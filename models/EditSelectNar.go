package models


type SelectNar struct {

	Id			string	`json:"id"`
	NumberNar		string	`json:"number_nar"`
	DateOpenNar		string	`json:"date_open_nar"`
	VrachOrtoped		string	`json:"vrach_ortoped"`
	DateStartProduction	string	`json:"date_start_production"`
	VrachTechnic		string	`json:"vrach_technic"`
	DateCloseNar		string	`json:"date_close_nar"`
	Sum			string	`json:"sum"`
}
