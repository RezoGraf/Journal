package main

import (
	"github.com/go-martini/martini"
	"./middleware/render"
	"./routes"
)


func main() {


	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))



	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Get("/reports", routes.Reports)
	m.Get("/login", routes.Login)
	m.Post("/logout", routes.LogOut)
	m.Get("/current_med_reg", routes.CurrentMedRegGet)
	m.Get("/", routes.Login)
	m.Post("/take_to_repair_comment_put", routes.TakeToRepairCommentPut)
	m.Get("/take_to_repair", routes.TakeToRepair)
	m.Post("/take_to_repair_patient_put", routes.TakeToRepairPatientPut)
	m.Post("/take_to_repair_patient_complite", routes.TakeToRepairPatientComplite)
	m.Post("/take_to_repair_invitation_put", routes.TakeToRepairInvitationPut)
	m.Get("/take_to_repair_queue_get", routes.TakeToRepairQueue)
	m.Get("/take_to_repair_invitation_get", routes.TakeToRepairInvitationGet)
	m.Get("/reports_invitation_get", routes.ReportsInvitationGet)
	m.Get("/take_to_repair_treatment_get", routes.TakeToRepairTreatmentGet)
	m.Get("/take_to_repair_treatment", routes.TakeToRepairTreatment)
	m.Get("/take_to_repair_treatment_get", routes.TakeToRepairTreatmentGet)
	m.Get("/taket_to_repair_get_info_patient", routes.TakeToRepairGetInfoPatient)
	m.Get("/taket_to_repair_get_info_patient_comments", routes.TakeToRepairGetInfoPatientCommentsGet)
	m.Post("/take_to_repair_move_patient_treatment", routes.TakeToRepairMovePatientTreatment)
	m.Get("/office_directory", routes.Office_directory)
	m.Get("/edit", routes.Edit)
	m.Post("/add_row", routes.Add_row)
	m.Get("/edit_get", routes.EditFormGet)
	m.Get("/edit_fiovrach_get", routes.EdtFioVrachGet)
	m.Get("/edit_fiovrach_technic_get", routes.EditFioVrachTechnicGet)
	m.Post("/edit_add_nar", routes.EditAddNar)
	m.Get("/edit_listNar_get", routes.EditListNarGet)
	m.Post("/edit_list_nar_edit_save", routes.EditListNarEditSave)
	m.Post("/save", routes.Save)
	m.Get("/catalog_vrach", routes.CatalogVrach)
	m.Get("/catalog_vrach_get", routes.CatalogVrachGet)
	m.Post("/catalog_vrach_update", routes.CatalogVrachUpdate)
	m.Post("/catalog_vrach_delete_row", routes.CatalogVrachDeleteRow)
	m.Post("/catalog_vrach_put", routes.CatalogVrachPut)
	m.Get("/catalog_thepost_get", routes.CatalogThePostGet)
	m.Get("/catalog_thepost", routes.CatalogThePost)
	m.Post("/add_the_post", routes.AddThePost)
	m.Post("/catalog_thepost_update", routes.CatalogThePostUpdate)
	m.Post("/catalog_thepost_delete_row", routes.CatalogThePostDeleteRow)
	m.Post("/addvrach", routes.AddVrach)
	m.Get("/catalog_med_reg_get", routes.CatalogMedRegGet)
	m.Post("/login_auth", routes.LoginAuth)
	m.Post("/addmedreg", routes.AddMedReg)
	m.Get("/catalog_lgot_cat_get", routes.CatalogLgotCatGet)
	m.Post("/addlgotcat", routes.AddLgotCat)
	m.Get("/listlgotcat", routes.LgotcatListView)
	m.Get("/find", routes.Find)
	m.Post("/findById", routes.FindById)
	m.Post("/FindByDateRecord", routes.FindByDateRecord)
	m.Post("/findByDateInvitation", routes.FindByDateInvitation)
	m.Post("/findByFio", routes.FindByFio)
	m.Get("/reportWeek", routes.ReportWeek)
	m.Post("/editPatientPhoneEdit", routes.EditPatientPhoneEdit)
	m.Get("/listNarGetPhone", routes.ListNarGetPhone)
	m.Post("/print_queue", routes.PrintQueue)
	m.Post("/print_report_week", routes.PrintReportWeek)
	m.Get("/take_to_repair_next_row", routes.TakeToRepairNextRow)
	m.Get("/take_to_repair_back_row", routes.TakeToRepairBacktRow)
	m.Get("/take_to_repair_treatment_next_row", routes.TakeToRepairNextRowTreatment)
	m.Get("/take_to_repair_treatment_back_row", routes.TakeToRepairBacktRowTreatment)
	m.Get("/take_to_repair_treatment_lgotcat", routes.TakeToRepairTreatmentLgotCat)
	m.Get("/patient_caused", routes.PatientCaued)
	m.RunOnAddr(":3000")
}
