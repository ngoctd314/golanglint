package main

import "database/sql"

func rowserrcheck() {
	db := sql.DB{}
	rows, err := db.Query("select id from tb") // Wrong case
	if err != nil {
		// handle error
	}
	for rows.Next() {
		// handle rows
	}
	// good
	// if rows.Err() != nil {
	// 	// handle err
	// }
}
