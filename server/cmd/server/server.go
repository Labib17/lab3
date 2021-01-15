package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"../../db"
	_ "github.com/go-sql-driver/mysql"
)


func new_connecting() (*sql.DB, error) {
	conn := &db.Connection{
		name:   "payment_system",
		user:     "helen",
		pin: "x7x",
		hosts:     "localhost:3306",
	}
	return conn.Open()
}

var dbCon, err = new_connecting()

func rowsToString(rows *sql.Rows) string {
	result := ""
	col := make([]string, 0)
	col, err = rows.Columns()
	for i := 0; i < len(col); i++ {
		result += col[i] + "\t"
	}
	result += "\n"
	for rows.Next() {
		var id, balance, last_operation string
		rows.Scan(&id, &balance, &last_operation)
		result += fmt.Sprintf("%s\t%s\t%s\n", id, balance,last_operation)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}


func starting_the_server() error {
	log.Print("Server started on\n127.0.0.1:", PORT)
	http.HandleFunc("/fetch", fetchAccounts)
	http.HandleFunc("/transfer", transfer)
	return http.ListenAndServe(":15000", nil)
}

func fetchAccounts(rw http.ResponseWriter, r *http.Request) {
	rows, err := db.FetchAccountsdb(dbCon)

	if err != nil {
		log.Fatal(err)
	}

	masterData := rowsToString(rows)
	rw.Write([]byte(masterData))
}

func id_of_account(rw http.ResponseWriter, r *http.Request) {
	amount, ok1 := r.URL.Query()["amount"]
	sender, ok2 := r.URL.Query()["sender"]
	receiver, ok3 := r.URL.Query()["receiver"]
	if !ok1 || !ok2 || !ok3 {
		log.Fatal("not ok")
	}

	date := time.Now().Format("2006.01.02 15:04:05")

	err := db.transfer(sender[0], recepient[0], amount[0], date, dbCon)

	if err != nil {
		rw.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	rw.Write([]byte("Money transfered successfully"))
}
