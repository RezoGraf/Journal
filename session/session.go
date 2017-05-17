package session

import (
	"encoding/gob"
	"net/http"
	"../framework/sessions"
)

var cookieStore = sessions.NewCookieStore([]byte("secret"))

const cookieName = "MyCookie"

type sesKey int

const (
	sesKeyLogin sesKey = iota
)


func SetAutch(w http.ResponseWriter ,r *http.Request , id, fam, name, lastname string){
	gob.Register(sesKey(0))
	ses, err := cookieStore.Get(r, cookieName)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	ses.Values[sesKeyLogin] = fam +" "+ name +" "+lastname
	err = cookieStore.Save(r, w, ses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CurrentUser(r *http.Request) string{
	ses, _ := cookieStore.Get(r, cookieName)
	login, ok := ses.Values[sesKeyLogin].(string)
	if !ok {
		return "fail"
	} else {
		return login
	}

}

func ClearSession(r *http.Request, w http.ResponseWriter) {

	ses, _ := cookieStore.Get(r, cookieName)
	for key, _ := range ses.Values {
		delete(ses.Values, key)
	}
	ses.Save(r, w)
}

func GetAutch(w http.ResponseWriter, r *http.Request,) (string){
	ses, _ := cookieStore.Get(r, cookieName)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return "error"
	//}

	_ , ok := ses.Values[sesKeyLogin].(string)
	if ok {
		//login = "anonymous"
		return "ok"
	} else {
		return "fail"
	}
}
