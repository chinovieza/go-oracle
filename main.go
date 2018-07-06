package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func main() {

	//https://gist.github.com/mnadel/8678269
	//but don't use 64-bit tdm-gcc
	//solve by below
	//https://github.com/mattn/go-oci8/issues/75
	//#oqamar
	//https://gocodecloud.com/blog/2016/08/09/accessing-an-oracle-db-in-go/

	db, err := sql.Open("oci8", "username/password@localhost:1521/xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

	rows, err := db.Query("SELECT VALUE FROM DMPSS_TEMP WHERE KEY = 'ENV'")
	if err != nil {
		fmt.Println("Error fetching addition")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var v string
		rows.Scan(&v)
		fmt.Printf("Hello Oracle -> : %v\n", v)
	}
}
