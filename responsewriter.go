package main

import (
	"database/sql"
	"fmt"
)

func handleResponses() {
	fmt.Println("Fetching Responses")
	responses := fetchResponses()
	for _, v := range responses {
		fmt.Println(v)
	}
}

func fetchResponses() []Response {
	db, err := sql.Open(driverName, dataSourceName)
	check(err)
	defer db.Close()
	query := "SELECT res.id, res.user, res.response_note, res.response_date FROM reqres.responses res INNER JOIN reqres.requests req ON res.request_id = req.id"
	rows, err := db.Query(query)
	check(err)
	var id int
	var user, responseNote, responseDate string
	var responses []Response
	for rows.Next() {
		err = rows.Scan(&id, &user, &responseNote, &responseDate)
		check(err)
		responses = append(responses, Response{id, user, responseNote, responseDate})
	}
	return responses
}

type Response struct {
	id           int
	user         string
	responseNote string
	responseDate string
}
