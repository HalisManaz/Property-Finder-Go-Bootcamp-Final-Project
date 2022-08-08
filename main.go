package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome PF Market\n"+
		"If there are more than 3 items of the same product, then fourth and subsequent ones would have 8 percent off.\n"+
		"If you made purchase which is more than 500 in a month then all subsequent purchaseshave 10 percen off.\n"+
		"Every fourth order whose total is more than 100 may have discount depending on products which are discount products of the month. "+
		"Products whose VAT is 1 percent don’t have any discount but products whose VAT is 8 percent and 18 percent have discount of 10 percent and 15 percent respectively.\n"+
		"Only one discount can be applied at a time and only disconunt which is the highest are applied")
	if err != nil {
		return
	}
}

func ConnectSQL(sqldb string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:sqlpassword@tcp(127.0.0.1:3306)/"+sqldb)
	//defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
}
