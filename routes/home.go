package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"../utils"
	"../middleware/render"
	"../models"
	_ "github.com/lib/pq"
	//"../framework/sessions"
	//"encoding/gob"
	"../session"
)

//var cookieStore = sessions.NewCookieStore([]byte("secret"))
//
//const cookieName = "MyCookie"
//
//type sesKey int
//
//const (
//	sesKeyLogin sesKey = iota
//)

func Login(w http.ResponseWriter, r *http.Request, rnd render.Render)  {
	login := session.GetAutch(w, r)
	if login == "ok" {
		rnd.HTML(200, "take_to_repair", nil)
	} else {
		rnd.HTML(200, "login", nil)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request, rnd render.Render)  {
	session.ClearSession(r, w)
	rnd.HTML(200, "/login", nil)
}
func Reports(w http.ResponseWriter, r *http.Request, rnd render.Render)  {
	value := session.GetAutch(w,r)
	if value == "ok" {
		rnd.HTML(200, "reports", nil)
	} else {
		w.Write([]byte("Не авторизованный пользователь"))
	}
}


func TakeToRepair(w http.ResponseWriter, r *http.Request, rnd render.Render) { //Страница добавление пациента
	value := session.GetAutch(w,r)
	//ses, err := cookieStore.Get(r, cookieName)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//_ , ok := ses.Values[sesKeyLogin].(string)
	if value == "ok" {
		//login = "anonymous"
		rnd.HTML(200, "take_to_repair", nil)
	} else {
		w.Write([]byte("Не авторизованный пользователь"))
	}

	//w.Write([]byte("you are " + login))
	//rnd.HTML(200, "take_to_repair", nil)
}


type currentIdListNar struct {
	id string
}

func PrintReportWeek(rnd render.Render, r *http.Request) { //Печать очереди
	bks := models.ModelPrintReportWeek()
	rnd.HTML(200, "print_report_week", bks)
}

func PrintQueue(rnd render.Render, r *http.Request) { //Печать очереди
	r.ParseForm()
	bks := models.ModelPrintQueue(r.FormValue("id"))
	rnd.HTML(200, "print_queue", bks)
}

var currentId currentIdListNar //Экземпляр необходим для функций LinstNar и Edit

func EditListNarGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bks := models.EditModelListNarGet(r.FormValue("id"))
	b, err := json.Marshal(bks)

	if err != nil{
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func ListNarGetPhone(w http.ResponseWriter) {
	bks := models.ModelListNarGetPhone(currentId.id)
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)

}

//noinspection ALL
func EditListNarEditSave(rnd render.Render, r *http.Request) {
	r.ParseForm()
	models.ModelEditNarEditSave(r.FormValue("id"),
		r.FormValue("id_patient"),
		r.FormValue("date_open_nar"),
		r.FormValue("vrach_ortoped"),
		r.FormValue("date_start_production"),
		r.FormValue("vrach_technic"),
		r.FormValue("date_close_nar"),
		r.FormValue("sum"),
		r.FormValue("name_reg"))

}

func CatalogThePostUpdate(r *http.Request) {
	r.ParseForm()
	models.ModelCatalogThePostUpdate(r.FormValue("id"), r.FormValue("the_post"))
}

func CatalogVrachUpdate(r *http.Request) {
	r.ParseForm()
	models.ModelCatalogVrachUpdate(r.FormValue("id"), r.FormValue("fam"), r.FormValue("name"), r.FormValue("lastname"), r.FormValue("the_post"))
}

func CatalogVrachPut(r *http.Request) {
	r.ParseForm()
	models.ModelCatalogVrachPut(r.FormValue("fam"), r.FormValue("name"), r.FormValue("lastname"), r.FormValue("the_post"))
}

func CatalogVrachDeleteRow(r *http.Request) {
	r.ParseForm()
	models.ModelCatalogVrachDeleteRow(r.FormValue("id"))
}

func CatalogThePostDeleteRow(r *http.Request) {
	r.ParseForm()
	models.ModelCatalogThePostDeleteRow(r.FormValue("id"))
}

func EditPatientPhoneEdit(r *http.Request) {
	r.ParseForm()
	//fmt.Println(r.FormValue("id"), r.FormValue("patient_id"), r.FormValue("CloseOpen"))
	models.ModelEditPatientPhoneEdit(r.FormValue("id"), r.FormValue("phone"))
}

func FormatDate(Time time.Time) string {
	return Time.Format("2006-01-02")
}

func EditAddNar(rnd render.Render, r *http.Request) {
	r.ParseForm()
	id_patient := r.FormValue("id_patient")
	vrach_ortoped := r.FormValue("vrach_ortoped")
	vrach_technic := r.FormValue("vrach_technic")
	number_nar := r.FormValue("number_nar")
	date_open_nar := r.FormValue("date_open_nar")
	date_start_production := r.FormValue("date_start_production")
	date_close_nar := r.FormValue("date_close_nar")
	sum := r.FormValue("sum")
	name_reg := r.FormValue("name_reg")

	fmt.Println("name_reg = "+name_reg)

	models.ModelAddNar(`INSERT INTO j_nar (
        id_patient,
        vrach_ortoped,
        vrach_technic,
        number_nar,
        date_open_nar,
        date_start_production,
        date_close_nar,
        sum,
        created
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		id_patient,
		utils.NullableString(vrach_ortoped),
		utils.NullableString(vrach_technic),
		utils.NullableInt(number_nar),
		utils.NullableString(date_open_nar),
		utils.NullableString(date_start_production),
		utils.NullableString(date_close_nar),
		utils.NullableFloat(sum),
		utils.NullableString(name_reg))
}

func FindByFio(rnd render.Render, r *http.Request) {
	r.ParseForm()
	rnd.HTML(200, "find", nil)
}

func FindByDateInvitation(rnd render.Render, r *http.Request) { //Поиск пациента по дате приглашения
	r.ParseForm()
	bks := models.ModelFindByDateInvitation(r.FormValue("date_invitation"))

	rnd.HTML(200, "find", bks)
}

func FindByDateRecord(rnd render.Render, r *http.Request) { //Поиск пациента по дате записи
	r.ParseForm()
	bks := models.ModelFindByDateRecord(r.FormValue("date_record"))
	rnd.HTML(200, "find", bks)
}

func FindById(rnd render.Render, r *http.Request) { //Поиск пциента по id
	r.ParseForm()
	bks := models.ModelFindById(r.FormValue("id"))
	rnd.HTML(200, "find", bks)
}

func Find(rnd render.Render) {
	rnd.HTML(200, "find", nil)
}

func Edit(rnd render.Render) {
	rnd.HTML(200, "edit", nil)
}

func IndexHandler(rnd render.Render) {
	rnd.HTML(200, "index", nil)
}

func CatalogLgotCatGet(w http.ResponseWriter, rnd render.Render) { //Просмотр льготных категорий
	bks := models.ModelCatalogLgotCatGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func TakeToRepairTreatmentLgotCat(rnd render.Render, r *http.Request) { //Просмотр льготных категорий
	r.ParseForm()
	bks := models.ModelTakeToRepairTreatmentLgotCat(r.FormValue("nameLong"))
	rnd.HTML(200, "take_to_repair_treatment", bks)
}

func TakeToRepairGetInfoPatient(r *http.Request, w http.ResponseWriter)  {
	r.ParseForm()
	bks := models.ModelTakeToRepairGetInfoPatient(r.FormValue("id"))
	b, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func AddLgotCat(w http.ResponseWriter, r *http.Request, rnd render.Render) {
	r.ParseForm()
	row := models.Add_rows{
		Name:     r.FormValue("name"),
		NameLong: r.FormValue("nameLong"),
	}
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/journal")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Prepare("INSERT INTO j_catalog_lgot_cat (name, name_long) VALUES ($1, $2)")
	rows.Exec(row.Name, row.NameLong)

	db.Close()

}

func CatalogMedRegGet(w http.ResponseWriter, rnd render.Render) { //Список мед регистраторов
	bks := models.ModelCatalogMedRegGet()
	b , err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func LoginAuth(w http.ResponseWriter, r *http.Request, rnd render.Render)  {
	r.ParseForm()
	bks, currentMedReg := models.ModelLoginAuth(r.FormValue("id"), r.FormValue("pwd"))
	if bks[0].Count == "1" {
		//--------------------------------------------------------
		//gob.Register(sesKey(0))
		//ses, err := cookieStore.Get(r, cookieName)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}
		//ses.Values[sesKeyLogin] = "user"
		//err = cookieStore.Save(r, w, ses)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}
		session.SetAutch(w,r, r.FormValue("id"), currentMedReg[0].Fam, currentMedReg[0].Name, currentMedReg[0].LastName)

		//--------------------------------------------------------
		rnd.Redirect("take_to_repair", 200)
	} else {
		rnd.HTML(403, "login", nil)
	}


}

func CurrentMedRegGet(w http.ResponseWriter, r *http.Request)  {
	user := session.CurrentUser(r)
	b , err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s \n", b)
	}
}


func AddMedReg(w http.ResponseWriter, r *http.Request, rnd render.Render) {
	r.ParseForm()
	row := models.Add_rows{
		Fam:      r.FormValue("fam"),
		Name:     r.FormValue("name"),
		Lastname: r.FormValue("lastname"),
	}
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/journal")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Prepare("INSERT INTO j_catalog_med_reg (fam, name, lastname) VALUES ($1, $2, $3)")
	rows.Exec(row.Fam, row.Name, row.Lastname)

	db.Close()

}

func AddVrach(w http.ResponseWriter, r *http.Request, rnd render.Render) {
	r.ParseForm()
	row := models.Add_rows{
		Fam:      r.FormValue("fam"),
		Name:     r.FormValue("name"),
		Lastname: r.FormValue("lastname"),
	}
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/journal")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Prepare("INSERT INTO j_catalog_vrach (fam, name, lastname) VALUES ($1, $2, $3)")
	rows.Exec(row.Fam, row.Name, row.Lastname)

	db.Close()

}

func CatalogVrach(w http.ResponseWriter, rnd render.Render) { //Список врачей
	rnd.HTML(200, "catalog_vrach", nil)
	//fmt.Fprintf(w, "%s \n", b)
}

func CatalogVrachGet(rnd render.Render, w http.ResponseWriter) {
	bks := models.ModelCatalogVrachGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func CatalogThePostGet(w http.ResponseWriter, rnd render.Render) {
	bks := models.ModelCatalogThePost()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
	//rnd.JSON(200, b)
}
func CatalogThePost(w http.ResponseWriter, rnd render.Render) {
	rnd.HTML(200, "catalog_thepost", nil)
}

func AddThePost(rnd render.Render, r *http.Request) {
	models.ModelAddThePost(r.FormValue("the_post"))
}

func Save(rnd render.Render, r *http.Request, w http.ResponseWriter) { //Сохранение EditForm
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 0, 64)
	fmt.Println(r.FormValue("flag_patient_complite"), "home.go")
	models.ModelSave(id,
		r.FormValue("id"),
		r.FormValue("phone"),
		r.FormValue("fiovrach"),
		r.FormValue("date_invitation"),
		r.FormValue("date_nar"),
		r.FormValue("numbernar"),
		r.FormValue("fioreg"),
		r.FormValue("comment"),
		r.FormValue("flag_patient_complite"),
		r.FormValue("flag_patient_refuse"),
	)
}

func EditFormGet(w http.ResponseWriter, rnd render.Render, r *http.Request) { //Редактирование пациента
	r.ParseForm()
	bks := models.ModelEditFormGet(r.FormValue("id"))
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func EdtFioVrachGet(w http.ResponseWriter) {
	bks := models.ModelEdtFioVrachGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func EditFioVrachTechnicGet(w http.ResponseWriter) {
	bks := models.ModelEditFioVrachTechnicGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}



func TakeToRepairQueue(w http.ResponseWriter)  {
	bks := models.ModelTakeToRepairQueue()
	b, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func TakeToRepairInvitationGet(w http.ResponseWriter)  {
	bks := models.ModelTakeToRepairInvitationGet()
	b, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func ReportsInvitationGet(w http.ResponseWriter, r *http.Request)  {
	bks := models.ModelReportsInvitationGet(r.FormValue("date_start"), r.FormValue("date_end"))
	b, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Fprintf(w, "%s \n", b)
	}
}
func take_to_repair_treatment_get(w http.ResponseWriter)  {
	bks := models.ModelTakeToRepairTreatmentGet()
	b, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Fprintf(w, "%s \n", b)
	}
}

func PatientCaued(rnd render.Render)  {
	rnd.HTML(200, "patient_caused", nil)
}

func TakeToRepairTreatment(rnd render.Render) {
	rnd.HTML(200, "take_to_repair_treatment", nil)
}

func TakeToRepairTreatmentGet(w http.ResponseWriter) {
	bks := models.ModelTakeToRepairTreatmentGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func TakeToRepairNextRow(w http.ResponseWriter, rnd render.Render, r *http.Request) { //Навигция вперед пациенты в очереди
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	idValue := id
	idValueString := strconv.FormatInt(idValue, 10)
	bks := models.ModelTakeToRepairNextRow(idValueString)
	rnd.HTML(200, "take_to_repair", bks)
}

func TakeToRepairBacktRow(w http.ResponseWriter, rnd render.Render, r *http.Request) { //Навигция назад пацинты в очереди
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	idValue := id
	idValueString := strconv.FormatInt(idValue, 10)
	bks := models.ModelTakeToRepairBacktRow(idValueString)
	rnd.HTML(200, "take_to_repair", bks)
}

//-------------------------Навигация пациенты на лечении---------------------------------------------
func TakeToRepairNextRowTreatment(w http.ResponseWriter, rnd render.Render, r *http.Request) { //Навигция вперед пациенты на лечении
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	idValue := id
	idValueString := strconv.FormatInt(idValue, 10)
	bks := models.ModelTakeToRepairNextRowTreatment(idValueString)
	rnd.HTML(200, "take_to_repair_treatment", bks)
}

func TakeToRepairBacktRowTreatment(w http.ResponseWriter, rnd render.Render, r *http.Request) { //Навигция назад пациенты на лечении
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	idValue := id
	idValueString := strconv.FormatInt(idValue, 10)
	bks := models.ModelTakeToRepairBacktRowTreatment(idValueString)
	rnd.HTML(200, "take_to_repair_treatment", bks)
}

//------------------------Конец навигации пациенты на лечении--------------------------------------
func LgotcatListView(w http.ResponseWriter, rnd render.Render) { //Просмотр льготных категорий
	//bks := models.ModelCatalogLgotCat()
	//for _, bk := range bks {
	//	fmt.Fprintf(w, "%s \n", bk.Name)
	//}

	bks := models.ModelCatalogLgotCatGet()
	b, _ := json.Marshal(bks)
	fmt.Fprintf(w, "%s \n", b)
}

func TakeToRepairPatientPut(r *http.Request)  {
	r.ParseForm()
	models.ModelTakeToRepairPatientPut(r.FormValue("fam"), r.FormValue("name"), r.FormValue("lastname"),
		r.FormValue("date_birth"), r.FormValue("number_phone"), r.FormValue("home_adres"),
		r.FormValue("numer_ud"), r.FormValue("number_pasport") ,r.FormValue("lgot_cat"), r.FormValue("fio_reg"),
		r.FormValue("note"))
}

func TakeToRepairInvitationPut(r *http.Request)  {
	r.ParseForm()
	models.ModelTakeToRepairInvitationPut(r.FormValue("id"), r.FormValue("date_invitation"), r.FormValue("comment"), r.FormValue("patient_refuse"))
}

func TakeToRepairPatientComplite(r *http.Request)  {
	models.ModelTakeToRepairPatientComplite(r.FormValue("id"))
}

func TakeToRepairMovePatientTreatment(r *http.Request)  {
	r.ParseForm()
	models.ModelTakeToRepairMovePatientTreatment(r.FormValue("id"))
}

func Add_row(w http.ResponseWriter, r *http.Request, rnd render.Render) {
	r.ParseForm()
	row := models.Add_rows{
		Fam:            r.FormValue("fam"),
		Name:           r.FormValue("name"),
		Lastname:       r.FormValue("lastname"),
		Datebirth:      r.FormValue("date_birth"),
		Phone:          r.FormValue("phone"),
		Homeadres:      r.FormValue("homeadres"),
		Numberud:       r.FormValue("numberud"),
		Lgotcat:        r.FormValue("lgotcat"),
		Fiovrach:       r.FormValue("fiovrach"),
		DateInvitation: r.FormValue("date_invitation"),
		DateNar:        r.FormValue("date_nar"),
		Numbernar:      r.FormValue("numbernar"),
		Fioreg:         r.FormValue("fioreg"),
		Comment:        r.FormValue("comment"),
	}
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/journal")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Prepare("INSERT INTO j_patient (fam, name, lastname, date_birth, phone, homeadres, numberud, lgotcat, fiovrach, date_invitation, date_nar, numbernar, fioreg, comment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)")
	rows.Exec(row.Fam, row.Name, row.Lastname, row.Datebirth, row.Phone, row.Homeadres, row.Numberud, row.Lgotcat, row.Fiovrach, row.DateInvitation, row.DateNar, row.Numbernar, row.Fioreg, row.Comment)

	db.Close()

}
func Office_directory(rnd render.Render) {
	rnd.HTML(200, "office_directory", nil)
}

func ReportWeek(rnd render.Render) {
	bks := models.ModelReportWeek()
	rnd.HTML(200, "reportWeek", bks)
}
