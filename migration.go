package main

import (
	"io/ioutil"
	"log"
	"strings"
)

// ConfigDB variable
var ConfigDB DBConfig

// SQLString Variable
type SQLString string

// DBConfig struct
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// InitMigration function
func InitMigration(conf DBConfig) {
	ConfigDB = conf
}

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

		requests = strings.Split(string(file), ";")

	}

	sqlStrings := make([]SQLString, len(requests))

	for index, request := range requests {
		sqlStrings[index] = SQLString(request)
		// result, err := db.Exec(request)
		// do whatever you need with result and error
	}

	return &sqlStrings
}

// RunSQL function
func RunSQL(sqlStrings *[]SQLString) {
	// for index, string := range *sqlStrings {
	// result, err := db.Exec(request)
	// do whatever you need with result and error
	// }
}
