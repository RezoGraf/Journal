package models

import (
	"strconv"
	"time"

	"../db"
	"../utils"
	//"github.com/fatih/motion/astcontext"
	//"fmt"
	"fmt"
	//"crypto/tls"
	"database/sql"
	//"database/sql/driver"
	"github.com/lib/pq"
	//"net/http"
	//"../session"
)



type PatientEditView struct {
	Id                  string  `json:"Id"`
	Fam                 string `json:"Fam"`
	Name                string `json:"Name"`
	Lastname            string `json:"Lasname"`
	Datebirth           string `json:"Datebirth"`
	Pasport		    string `json:"Pasport"`
	Phone               string `json:"Phone"`
	Homeadres           string `json:"Homeadres"`
	Numberud            string `json:"Numberud"`
	Lgotcat             string `json:"Lgotcat"`
	DateInvitation      string `json:"DateInvitation"`
	Fioreg              string `json:"Fioreg"`
	Comment             string `json:"Comment"`
	FlagPatientRefuse   string `json:"FlagPatientRefuse"`
	FlagPatientComplite string `json:"FlagPatientComplite"`
	FullName	    string `json:"FullName"`
	DateRecord	    string `json:"DateRecord"`
	Num	  	    string `json"Num`
}

type PatientCommentsView struct {
	PatientId string `json"Id"`
	Comment string `json:"Comment"`
	DateComment string `json:"DateComment"`
	NameReg string `json:"NameReg"`
}

type Add_rows struct {
	Id        int64
	IdPatient string `json:"IdPatient"`
	Fam       string
	Name      string `json:"Name"`
	NameLong  string //addLgotcat
	Lastname  string
	Datebirth string
	Phone     string
	Homeadres string
	Numberud  string
	Lgotcat   string
	// Lgotcat_1[]	 string
	Fiovrach       string
	DateInvitation string
	DateNar        string
	Numbernar      string `json:"Numbernar"`
	Fioreg         string
	Comment        string
	NumberNar      string    `json:"NumberNar"`
	DateOpenNar    time.Time `json:"DateOpen"`
	DateCloseNar   time.Time `json:"CloseOpen"`
}

type ListNar struct {
	Id           int64  `json:"Id"`
	IdPatient    int64  `json:"IdPatient"`
	NumberNar    string `json:"NumberNar"`
	DateOpenNar  string `json:"DateOpen"`
	DateCloseNar string `json:"CloseOpen"`
	Sum          string `json:"Sum"`
}

type Phone struct {
	Id    int64  `json:"Id"`
	Phone string `json:"Phone"`
}
type CatalogVrach struct {
	Id       int64  `json:"Id"`
	Fam      string `json:"Fam"`
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	ThePost  string `json:"ThePost"`
}

type CatalogPost struct {
	Id      int64  `json:"Id"`
	ThePost string `json:"ThePost"`
}

//func ModelReportWeek()[]*Table_view{
//	query := "SELECT id, fam, name, lastname, date_birth FROM j_patient WHERE date_invitation = '" + date_invitation + "'"
//	rows := db.Select(query)
//	bks := make([]*Table_view, 0)
//	for rows.Next() {
//		bk := new(Table_view)
//		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth)
//		bks = append(bks, bk)
//	}
//	return bks
//}

func ModelFindByDateInvitation(date_invitation string) []*Table_view { //Поиск пациента по дате приглашения
	query := "SELECT id, fam, name, lastname, date_birth FROM j_patient WHERE date_invitation = '" + date_invitation + "'"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth)
		bks = append(bks, bk)
	}
	return bks
}

func ModelFindByDateRecord(date_record string) []*Table_view { //Поиск пациента по дате записи
	query := "SELECT id, fam, name, lastname, date_birth FROM j_patient WHERE date_record = '" + date_record + "'"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth)
		bks = append(bks, bk)
	}
	return bks
}

func ModelFindById(id string) []*Table_view { //Поиск пциента по id
	query := "SELECT id, fam, name, lastname, date_birth FROM j_patient WHERE id = " + id + ""
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth)
		bks = append(bks, bk)
	}
	return bks
}

func ModelPrintReportWeek() []*Table_view { //Печать пролеченых пациентов за неделю
	query := "SELECT EXTRACT(WEEK FROM TIMESTAMP 'now') id_d, id, fam, name, lastname, date_birth, date_last_edit FROM j_patient where flag_patient_complite = true"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id_d, &bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateLastEdit)
		bks = append(bks, bk)
	}
	return bks
}
func ModelPrintQueue(id string) []*Table_view { //Печать очереди
	query := "SELECT id, fam, name, lastname, date_birth, numberud, date_record, phone,lgotcat FROM j_patient WHERE id = '" + id + "'"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.Numberud, &bk.DateRecord, &bk.Phone, &bk.Lgotcat)
		bks = append(bks, bk)
	}
	return bks
}

func EditModelListNarGet(id string) []*SelectNar {
	rows := db.Select(`
	SELECT
	id, number_nar, date_open_nar, vrach_ortoped, date_start_production, vrach_technic, date_close_nar, sum
	FROM j_nar
	WHERE id_patient = $1`, id)
	var NumberNarnullStr sql.NullString
	var DateOpenNarNullStr pq.NullTime
	var VrachOrtopedNullStr sql.NullString
	var DateStartProductionNullStr pq.NullTime
	var VrachTechnicNullstr sql.NullString
	var DateCloseNarNullStr pq.NullTime
	var SummNulFloat64 sql.NullString
	bks := make([]*SelectNar, 0)
	for rows.Next() {
		bk := new(SelectNar)
		rows.Scan(&bk.Id, &NumberNarnullStr,  &DateOpenNarNullStr, &VrachOrtopedNullStr, &DateStartProductionNullStr, &VrachTechnicNullstr, &DateCloseNarNullStr, &SummNulFloat64)
		bk.NumberNar = NumberNarnullStr.String
		bk.DateOpenNar = utils.FormatDatePqNullTime(DateOpenNarNullStr)
		bk.VrachOrtoped = VrachOrtopedNullStr.String
		bk.DateStartProduction = utils.FormatDatePqNullTime(DateStartProductionNullStr)
		bk.VrachTechnic = VrachTechnicNullstr.String
		bk.DateCloseNar = utils.FormatDatePqNullTime(DateCloseNarNullStr)
		bk.Sum = SummNulFloat64.String
		bks = append(bks, bk)
	}
	return bks
}


func ModelListNarGetPhone(id string) []*Phone {
	query := "SELECT id, phone from j_patient where id = " + id + ""
	rows := db.Select(query)
	bks := make([]*Phone, 0)
	for rows.Next() {
		bk := new(Phone)
		rows.Scan(&bk.Id, &bk.Phone)
		bks = append(bks, bk)
	}
	return bks
}

func FormatDate(Time time.Time) string {
	return Time.Format("2006/01/02")
}

func ModelEditNarEditSave(id,
id_patient,
date_open_nar,
vrach_ortoped,
date_start_production,
vrach_technic,
date_close_nar,
sum,
name_reg string){
	db.Exec(`
	UPDATE j_nar SET date_open_nar = $1, vrach_ortoped = $2, date_start_production = $3, vrach_technic = $4, last_edit = $5,
		date_close_nar = $6, sum = $7
	WHERE id = $8 AND id_patient = $9
	`, utils.NullableTime(date_open_nar),
		utils.NullableString(vrach_ortoped),
		utils.NullableTime(date_start_production),
		utils.NullableString(vrach_technic),
		utils.NullableString(name_reg),
		utils.NullableTime(date_close_nar),
		utils.NullableFloat(sum),
		utils.NullableInt(id),
		utils.NullableInt(id_patient))


}

func ModelTakeToRepairInvitationPut(id, date_invitation, patient_refuse string)  {
	db.Exec(`UPDATE j_patient SET date_invitation = $1, flag_patient_refuse = $2
	WHERE id = $3
	`, utils.NullableString(date_invitation),
		utils.NullableString(patient_refuse),
		utils.NullableInt(id))
}

func ModelTakeToRepairCommentPut(id, comment, name_reg string)  {
	db.Exec(`INSERT INTO j_comments
	(id_patient, comment, name_reg)
	VALUES ($1, $2, $3)`,
	utils.NullableInt(id), utils.NullableString(comment), utils.NullableString(name_reg))
}

func ModelTakeToRepairPatientComplite(id string)  {
	db.Exec(`UPDATE j_patient SET flag_patient_complite = true
	WHERE id = $1
	`, utils.NullableString(id))
}

func ModelTakeToRepairMovePatientTreatment(id string)  {
	db.Exec(`UPDATE j_patient SET flag_on_treatment = true
	WHERE id = $1
	`, utils.NullableInt(id))
}

func ModelCatalogThePostUpdate(id, the_post string) {
	query := "update j_catalog_thepost SET name_the_post ='" + the_post + "' WHERE id = '" + id + "'"
	db.Update(query)
}

func ModelCatalogVrachUpdate(id, fam, name, lastname, the_post string) {
	query := "update j_catalog_vrach SET fam = '" + fam + "', name = '" + name + "', lastname = '" + lastname + "', name_the_post = '" + the_post + "' WHERE id = '" + id + "'"
	fmt.Println(query)
	db.Update(query)
}

func ModelCatalogVrachPut(fam, name, lastname, the_post string) {
	query := "insert into j_catalog_vrach (fam, name, lastname, name_the_post) values ('" + fam + "', '" + name + "', '" + lastname + "', '" + the_post + "')"
	db.Insert(query)
}

func ModelCatalogVrachDeleteRow(id string) {
	fmt.Println(id)
	query := "delete from j_catalog_vrach WHERE id = '" + id + "'"
	db.Delete(query)
}

func ModelCatalogThePostDeleteRow(id string) {
	fmt.Println(id)
	query := "delete from j_catalog_thepost WHERE id = '" + id + "'"
	db.Delete(query)
}

func ModelEditPatientPhoneEdit(id, phone string) {
	query := "UPDATE j_patient SET phone = '" + phone + "'  WHERE id = '" + id + "'"
	db.Update(query)
}

func ModelCatalogMedRegGet() []*MedReg { //Cписок мед регистраторов
	query := "SELECT id, fam, left(name,1), left(lastname, 1) FROM j_catalog_vrach where name_the_post = 'Регистратор' OR name_the_post = 'Администратор' ORDER BY id DESC"
	rows := db.Select(query)
	bks := make([]*MedReg, 0)
	for rows.Next() {
		bk := new(MedReg)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.LastName)
		bks = append(bks, bk)
	}
	return bks
}

func ModelCurrentMedRegGet(id string) []*MedReg { //Cписок мед регистраторов
	query := "SELECT id, fam, left(name,1), left(lastname, 1) FROM j_catalog_vrach where id = $1"
	rows := db.Select(query, utils.NullableInt(id))
	bks := make([]*MedReg, 0)
	for rows.Next() {
		bk := new(MedReg)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.LastName)
		bks = append(bks, bk)
	}
	return bks
}



func ModelLoginAuth(id, pwd string) ([]*MedReg, []*MedReg) { //Cписок мед регистраторов
	rows := db.Select(`SELECT count(*)
	FROM j_catalog_vrach
	WHERE id = $1 AND pwd = $2
	`, utils.NullableInt(id), utils.NullableString(pwd))
	defer rows.Close()
	bks := make([]*MedReg, 0)
	for rows.Next() {
		bk := new(MedReg)
		var countStr sql.NullString
		rows.Scan(&countStr)
		bk.Count = countStr.String
		bks = append(bks, bk)
	}
	currentMedReg := ModelCurrentMedRegGet(id)
	return bks, currentMedReg

}

func ModelCatalogLgotCatGet() []*LgotCat { //Просмотр льготных категорий
	query := "SELECT id, name, name_long FROM j_catalog_lgot_cat"
	rows := db.Select(query)
	defer rows.Close()
	bks := make([]*LgotCat, 0)
	for rows.Next() {
		bk := new(LgotCat)
		rows.Scan(&bk.Id, &bk.Name, &bk.NameLong)
		bks = append(bks, bk)
	}
	return bks
}

func ModelTakeToRepairGetInfoPatient(rId, type_query string)[]*PatientEditView  {
	var query = ""
	if type_query == "profile_patient"{
		query = `SELECT
	id, full_name, date_birth, number_pasport, date_record, homeadres, lgotcat, phone,
	(select count(*) from j_patient where id < $1 AND date_invitation IS NULL AND flag_patient_refuse = false)
	as num
	FROM j_patient
	WHERE id = $2`
	}else {

	query = `SELECT
	id, full_name, date_birth, number_pasport, date_record, homeadres, lgotcat, phone,
	(select count(*) from j_patient where id < $1)
	as num
	FROM j_patient
	WHERE id = $2`
	}
	rows := db.Select(query, rId, rId)
	var id sql.NullString
	var full_name sql.NullString
	var datebirth pq.NullTime
	var pasport sql.NullString
	var daterecord pq.NullTime
	var homeadres sql.NullString
	var lgotcat sql.NullString
	var phone sql.NullString
	var num sql.NullString
	bks := make([]*PatientEditView, 0)
	for rows.Next() {
		bk := new(PatientEditView)
		rows.Scan(&id, &full_name, &datebirth, &pasport, &daterecord, &homeadres, &lgotcat, &phone, &num)
		bk.Id = id.String
		bk.FullName = full_name.String
		bk.Datebirth = utils.FormatDatePqNullTime(datebirth)
		bk.Pasport = pasport.String
		bk.DateRecord = utils.FormatDatePqNullTime(daterecord)
		bk.Homeadres = homeadres.String
		bk.Lgotcat = lgotcat.String
		bk.Phone = phone.String
		bk.Num = num.String
		bks = append(bks, bk)
	}
	fmt.Print(type_query)
	return bks
}

func ModelTakeToRepairGetInfoPatientCommentsGet(rId string)[]*PatientCommentsView {
	rows := db.Select(`SELECT
	id_patient, comment, date_comment, name_reg FROM j_comments
	WHERE id_patient = $1`, rId)
	var id sql.NullString
	var comment sql.NullString
	var date_coment pq.NullTime
	var name_reg sql.NullString
	bks := make([]*PatientCommentsView, 0)
	for rows.Next() {
		bk := new(PatientCommentsView)
		rows.Scan(&id, &comment, &date_coment, &name_reg)
		bk.PatientId = id.String
		bk.Comment = comment.String
		bk.DateComment = utils.FormatDatePqNullTime(date_coment)
		bk.NameReg = name_reg.String
		bks = append(bks, bk)
	}
	return bks

}

func ModelCatalogVrachGet() []*CatalogVrach { //Список врачей
	query := "SELECT id, fam, name, lastname, name_the_post FROM j_catalog_vrach ORDER BY id DESC"
	rows := db.Select(query)
	bks := make([]*CatalogVrach, 0)
	for rows.Next() {
		bk := new(CatalogVrach)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.ThePost)
		bks = append(bks, bk)
	}
	return bks
}

func ModelCatalogThePost() []*CatalogPost {
	query := "SELECT id, name_the_post FROM j_catalog_thepost ORDER BY id DESC"
	rows := db.Select(query)
	bks := make([]*CatalogPost, 0)
	for rows.Next() {
		bk := new(CatalogPost)
		rows.Scan(&bk.Id, &bk.ThePost)
		bks = append(bks, bk)
	}
	return bks
}

func ModelAddThePost(the_post string) {
	fmt.Println("modelAddThePost " + the_post)
	query := "insert into j_catalog_thepost (name_the_post) values ('" + the_post + "')"
	db.Insert(query)

}

//-------------------------Навигация пациенты в очереди--------------------------------------------
func ModelTakeToRepairBacktRow(id string) []*Table_view { //Навигция назад
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is null and flag_patient_refuse = false and flag_patient_complite = false and j_patient.id < " + id + " ORDER BY j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord)
		bks = append(bks, bk)
	}
	return bks
}

func ModelTakeToRepairNextRow(id string) []*Table_view { //Навигция вперед
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is null and flag_patient_refuse = false and flag_patient_complite = false and j_patient.id > " + id + " ORDER BY j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord)
		bks = append(bks, bk)
	}
	return bks
}

//------------------------Конец навигации пациенты в очереди--------------------------------------
//-------------------------Навигация пациенты на лечении---------------------------------------------
func ModelTakeToRepairBacktRowTreatment(id string) []*Table_view { //Навигция назад п
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is not null and flag_patient_refuse = false and flag_patient_complite = false and j_patient.id < " + id + " ORDER BY j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord)
		bks = append(bks, bk)
	}
	return bks
}

func ModelTakeToRepairNextRowTreatment(id string) []*Table_view { //Навигция вперед
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is not null and flag_patient_refuse = false and flag_patient_complite = false and j_patient.id > " + id + " ORDER BY j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord)
		bks = append(bks, bk)
	}
	return bks
}

//------------------------Конец навигации пациенты на лечении--------------------------------------
func ModelTakeToRepair() []*Table_view { //Страница добавление пациента
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is null and flag_patient_refuse = false and flag_patient_complite = false order by j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord)
		bks = append(bks, bk)
	}
	return bks
}

//-----------------------Пациенты на лечении по льготным категориям----------------------------------
func ModelTakeToRepairTreatmentLgotCat(nameLong string) []*Table_view { //Страница добавление пациента
	query := "select distinct j_patient.id, fam, name, lastname, date_birth, date_record, lgotcat from j_patient left join j_nar on j_nar.id_patient = j_patient.id where date_open_nar is not null and flag_patient_refuse = false and flag_patient_complite = false and lgotcat = '" + nameLong + "' order by j_patient.id LIMIT 10"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.DateRecord, &bk.Lgotcat)
		bks = append(bks, bk)
	}
	return bks
}

//-----------------------Конец пациенты на лечении по льготным категориям----------------------------

//func ModelTakeToRepairTreatmentGet() []*PatientTreatment { //Страница добавление пациента
//	rows := db.Select(`SELECT
//	id, fam, name, lastname, date_birth, date_record, lgotcat
//	FROM j_patient
//	WHERE date_invitation IS NOT NULL`)
//	var id sql.NullString
//	var fam sql.NullString
//	var name sql.NullString
//	var lastname sql.NullString
//	var datebirth sql.NullString
//	var daterecord sql.NullString
//	var lgotcat sql.NullString
//	bks := make([]*PatientTreatment, 0)
//	for rows.Next() {
//		bk := new(PatientTreatment)
//		rows.Scan(&id, &fam, &name, &lastname, &datebirth, &daterecord, &lgotcat)
//		bk.Id.String = id.String
//		bk.Fam.String = fam.String
//		bk.Name.String = name.String
//		bk.Lastname.String = lastname.String
//		bk.DateRecord.String = daterecord.String
//		bk.Datebirth.String = datebirth.String
//		bk.Lgotcat.String = lgotcat.String
//		bks = append(bks, bk)
//	}
//	return bks
//}

func ModelEditFormGet(id string) []*PatientEditView { //Редактирование пациента
	query := "SELECT j_patient.id, fam, name, lastname, date_birth, phone, homeadres, numberud, lgotcat, date_invitation, fioreg, comment, flag_patient_refuse, flag_patient_complite FROM j_patient WHERE j_patient.id = " + id + ""
	rows := db.Select(query)
	bks := make([]*PatientEditView, 0)
	for rows.Next() {
		bk := new(PatientEditView)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.Datebirth, &bk.Phone, &bk.Homeadres, &bk.Numberud, &bk.Lgotcat, &bk.DateInvitation, &bk.Fioreg, &bk.Comment, &bk.FlagPatientRefuse, &bk.FlagPatientComplite)
		bks = append(bks, bk)
	}
	return bks
}

func ModelEdtFioVrachGet() []*CatalogVrach {
	query := "SELECT id, fam, left(name,1), left(lastname,1) from j_catalog_vrach where name_the_post = 'Врач-стоматолог-ортопед'"
	rows := db.Select(query)
	bks := make([]*CatalogVrach, 0)
	for rows.Next() {
		bk := new(CatalogVrach)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname)
		bks = append(bks, bk)
	}
	return bks
}

func ModelTakeToRepairQueue()[]*PatientTreatmentSrting  {
	rows := db.Select(`SELECT
	id, fam, name, lastname, full_name, date_birth, date_record, lgotcat
	FROM j_patient
	WHERE date_invitation IS NULL AND flag_patient_refuse = false order by id`)
	var id sql.NullString
	var fam sql.NullString
	var name sql.NullString
	var lastname sql.NullString
	var full_name sql.NullString
	var datebirth pq.NullTime
	var daterecord pq.NullTime
	var lgotcat sql.NullString
	bks := make([]*PatientTreatmentSrting, 0)
	for rows.Next() {
		bk := new(PatientTreatmentSrting)
		rows.Scan(&id, &fam, &name, &lastname, &full_name, &datebirth, &daterecord, &lgotcat)
		bk.Id = id.String
		bk.Fam = fam.String
		bk.Name = name.String
		bk.Lastname = lastname.String
		bk.FullName = full_name.String
		bk.DateRecord = utils.FormatDatePqNullTime(daterecord)
		bk.Datebirth = utils.FormatDatePqNullTime(datebirth)
		bk.Lgotcat = lgotcat.String
		bks = append(bks, bk)
	}
	return bks
}

func ModelTakeToRepairInvitationGet()[]*PatientTreatmentSrting  {
	rows := db.Select(`SELECT
	id, fam, name, lastname, full_name, date_birth, date_record, lgotcat, date_invitation
	FROM j_patient
	WHERE date_invitation IS NOT NULL AND flag_patient_refuse = false AND flag_on_treatment = false`)
	var id sql.NullString
	var fam sql.NullString
	var name sql.NullString
	var lastname sql.NullString
	var full_name sql.NullString
	var datebirth pq.NullTime
	var daterecord pq.NullTime
	var lgotcat sql.NullString
	var date_invitation pq.NullTime
	bks := make([]*PatientTreatmentSrting, 0)
	for rows.Next() {
		bk := new(PatientTreatmentSrting)
		rows.Scan(&id, &fam, &name, &lastname, &full_name, &datebirth, &daterecord, &lgotcat, &date_invitation)
		bk.Id = id.String
		bk.Fam = fam.String
		bk.Name = name.String
		bk.Lastname = lastname.String
		bk.FullName = full_name.String
		bk.DateRecord = utils.FormatDatePqNullTime(daterecord)
		bk.Datebirth = utils.FormatDatePqNullTime(datebirth)
		bk.Lgotcat = lgotcat.String
		bk.DateInvitation = utils.FormatDatePqNullTime(date_invitation)
		bks = append(bks, bk)
	}
	return bks

}

func ModelReportsInvitationGet(date_start, date_end string)[]*PatientTreatmentSrting  {
	rows := db.Select(`SELECT
	id, fam, name, lastname, full_name, date_birth, date_record, lgotcat, date_invitation
	FROM j_patient
	WHERE date_invitation IS NOT NULL AND flag_patient_refuse = false AND flag_on_treatment = false
	AND date_invitation between $1 AND $2`, utils.NullableString(date_start), utils.NullableString(date_end))
	var id sql.NullString
	var fam sql.NullString
	var name sql.NullString
	var lastname sql.NullString
	var full_name sql.NullString
	var datebirth pq.NullTime
	var daterecord pq.NullTime
	var lgotcat sql.NullString
	var date_invitation pq.NullTime
	bks := make([]*PatientTreatmentSrting, 0)
	for rows.Next() {
		bk := new(PatientTreatmentSrting)
		rows.Scan(&id, &fam, &name, &lastname, &full_name, &datebirth, &daterecord, &lgotcat, &date_invitation)
		bk.Id = id.String
		bk.Fam = fam.String
		bk.Name = name.String
		bk.Lastname = lastname.String
		bk.FullName = full_name.String
		bk.DateRecord = utils.FormatDatePqNullTime(daterecord)
		bk.Datebirth = utils.FormatDatePqNullTime(datebirth)
		bk.Lgotcat = lgotcat.String
		bk.DateInvitation = utils.FormatDatePqNullTime(date_invitation)
		bks = append(bks, bk)
	}
	return bks

}

func ModelTakeToRepairTreatmentGet()[]*PatientTreatmentSrting  {
	rows := db.Select(`SELECT
	id, fam, name, lastname, full_name, date_birth, date_record, lgotcat
	FROM j_patient
	WHERE date_invitation IS NOT NULL AND flag_patient_refuse = false AND flag_on_treatment = true AND flag_patient_complite = false`)
	var id sql.NullString
	var fam sql.NullString
	var name sql.NullString
	var lastname sql.NullString
	var full_name sql.NullString
	var datebirth pq.NullTime
	var daterecord pq.NullTime
	var lgotcat sql.NullString
	bks := make([]*PatientTreatmentSrting, 0)
	for rows.Next() {
		bk := new(PatientTreatmentSrting)
		rows.Scan(&id, &fam, &name, &lastname, &full_name, &datebirth, &daterecord, &lgotcat)
		bk.Id = id.String
		bk.Fam = fam.String
		bk.Name = name.String
		bk.Lastname = lastname.String
		bk.FullName = full_name.String
		bk.DateRecord = utils.FormatDatePqNullTime(daterecord)
		bk.Datebirth = utils.FormatDatePqNullTime(datebirth)
		bk.Lgotcat = lgotcat.String
		bks = append(bks, bk)
	}
	return bks

}

func ModelEditFioVrachTechnicGet() []*CatalogVrach {
	query := "SELECT id, fam, left(name,1), left(lastname,1) from j_catalog_vrach where name_the_post = 'Зубной техник'"
	rows := db.Select(query)
	bks := make([]*CatalogVrach, 0)
	for rows.Next() {
		bk := new(CatalogVrach)
		rows.Scan(&bk.Id, &bk.Fam, &bk.Name, &bk.Lastname)
		bks = append(bks, bk)
	}
	return bks
}

func ModelAddNar(query, id_patient string, vrach_ortoped, vrach_technic sql.NullString, number_nar sql.NullInt64, date_open_nar ,date_start_production, date_close_nar sql.NullString, sum sql.NullFloat64, created sql.NullString)  {
	 db.Exec(query, id_patient, vrach_ortoped, vrach_technic, number_nar, date_open_nar ,date_start_production, date_close_nar, sum, created)
}

func ModelSave(id int64, idRow, phone, fiovrach, date_invitation, date_nar, number_nar, fioreg, comment, flag_patient_complite, flag_patient_refuse string) {
	flagPatientComplite, _ := strconv.ParseBool(flag_patient_complite)
	flagPatientRefuse, _ := strconv.ParseBool(flag_patient_refuse)
	row := Table_view{
		Id:                  id,
		Phone:               phone,
		Fiovrach:            fiovrach,
		DateInvitation:      date_invitation,
		DateNar:             date_nar,
		Numbernar:           number_nar,
		Fioreg:              fioreg,
		Comment:             comment,
		FlagPatientComplite: flagPatientComplite,
		FlagPatientRefuse:   flagPatientRefuse,
	}
	stringFlagPatientComplite := strconv.FormatBool(row.FlagPatientComplite)
	stringFlagPatientRefuse := strconv.FormatBool(row.FlagPatientRefuse)
	dateNow := utils.FormatDate(time.Now())
	query := "UPDATE j_patient SET phone = '" + row.Phone + "', fiovrach = '" + row.Fiovrach + "', date_invitation = '" + row.DateInvitation + "', date_nar = '" + row.DateNar + "', numbernar = '" + row.Numbernar + "', fioreg = '" + row.Fioreg + "', comment = '" + row.Comment + "', flag_patient_complite = " + stringFlagPatientComplite + ", flag_patient_refuse = " + stringFlagPatientRefuse + ", date_last_edit = '" + dateNow + "' WHERE id = '" + idRow + "' "
	db.Update(query)
}
func ModelReportWeek() []*Table_view {
	query := "SELECT EXTRACT(WEEK FROM TIMESTAMP 'now') id_d, id, fam, name, lastname, date_last_edit FROM j_patient where flag_patient_complite = true"
	rows := db.Select(query)
	bks := make([]*Table_view, 0)
	for rows.Next() {
		bk := new(Table_view)
		rows.Scan(&bk.Id_d, &bk.Id, &bk.Fam, &bk.Name, &bk.Lastname, &bk.DateLastEdit)
		bks = append(bks, bk)
	}
	return bks

}

func ModelTakeToRepairPatientPut(fam, name, lastname, date_birth, number_phone, home_adres, numer_ud, number_pasport, lgot_cat, fio_reg string)  {
	fullname := ""+fam+" "+name+" "+lastname+" "
	db.Exec(`INSERT INTO j_patient(
	fam, name, lastname, date_birth, phone, homeadres, numberud, number_pasport, lgotcat, fioreg, full_name)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
	utils.NullableString(fam),
	utils.NullableString(name),
		utils.NullableString(lastname),
		utils.NullableString(date_birth),
		utils.NullableString(number_phone),
		utils.NullableString(home_adres),
		utils.NullableString(numer_ud),
		utils.NullableString(number_pasport),
		utils.NullableString(lgot_cat),
		utils.NullableString(fio_reg),
		utils.NullableString(fullname))
}
