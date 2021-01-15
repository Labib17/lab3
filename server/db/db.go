package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type connecting struct {
	name         string
	user, pin    string
	hosts        string
}

func (c *connecting) kind_of_connecting() string {
	return fmt.Sprintf(c.name, c.user, c.pin, c.hosts)
}

func (c *connecting) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.kind_of_connecting())
}

func id_of_account(dbCon *sql.DB) (*sql.Rows, error) {
	rows, err := dbCon.Query("SELECT * FROM account;")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func transfer(sender, recepient, amount, date string, dbCon *sql.DB) error {
	var balance string
	err := dbCon.QueryRow("SELECT balance FROM accounts WHERE id = " + sender + ";").Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	an_amount, _ := strconv.ParseInt(amount, 10, 32)
	a_balance, _ := strconv.ParseInt(balance, 10, 32)

	if amountInt > balanceInt {
		return errors.New("Transfer failed")
	}
	
	sqlcom := "call transfer (" + sender + ", " + receiver + ", " + amount + ", '" + date + "');"
	_, err = dbCon.Exec(sqlcom)

	return nil
}


