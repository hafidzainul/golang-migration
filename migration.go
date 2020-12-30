package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// SQLString Variable
type SQLString string

// ReadSQL function
func ReadSQL(path string) *[]SQLString {

	var requests []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println("Reading file: " + f.Name())

		file, err := ioutil.ReadFile(path + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		reqs := strings.Split(string(file), ";\n")

		requests = append(requests, reqs...)

	}

	sqlStrings := make([]SQLString, len(requests))

	for index, request := range requests {
		sqlStrings[index] = SQLString(request)
	}

	return &sqlStrings
}

// RunSQL function
func RunSQL(sqlStrings *[]SQLString) {
	for index, query := range *sqlStrings {
		log.Println("Query (" + fmt.Sprint(index) + ") of " + fmt.Sprint(len(*sqlStrings)))
		log.Println("Executing query:", string(query))

		tx := Connection().Begin()

		if err := tx.Exec(string(query)).Error; err != nil {
			log.Fatal(err)
			tx.Rollback()
		} else {
			tx.Commit()
			log.Println("Successfully run query")
		}

	}
}
