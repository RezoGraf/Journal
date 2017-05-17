package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"fmt"
)

func db() (*sql.DB, error) {
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/journal")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, err
}

func Select(Query string, Args ...interface{}) *sql.Rows {
	db, _ := db()
	rows, err := db.Query(Query, Args...)
	if(err != nil){
		log.Fatal(err)
		db.Close()
	}else {
		return rows
		db.Close()

	}

	return rows
}

func Update(Query string){
	var err error
	db, _ := db()
	tx, _ := db.Begin()
	_, err = tx.Exec(Query)
	if err != nil {
		fmt.Println(err)
	} else{
		tx.Commit()
		db.Close()
	}

}
func Insert(Query string) {
	db, _ := db()
	tx, _ := db.Begin()
	_, _ = tx.Exec(Query)
	tx.Commit()
	db.Close()
}


func Exec(Query string, Args ...interface{}) {
	db, errdb := db()
	fmt.Println(Args...)
	fmt.Println(Query)
	if errdb !=nil {
		log.Fatal(errdb)
	}
	tx, txerr := db.Begin()
	if txerr == nil {
		_ , reserr := tx.Exec(Query, Args...)
		if reserr == nil {
			tx.Commit()
		} else {
			// Обрабатываем ошибку команды reserr
			tx.Rollback()
			fmt.Println(reserr)

		}
	} else {
		fmt.Println(txerr)
		// Тут обрабатываем ошибку создания транзакции txerr
	}
	db.Close()
}



func Delete(Query string) {
	db, _ := db()
	tx, _ := db.Begin()
	_, _ = tx.Exec(Query)
	tx.Commit()
	db.Close()
}

