package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strings"
)

type Status int

const (
	OK Status = iota
	InvalidArgs
	FileNotFound
	DatabaseFail
	OutputWrite
)

const (
	relativePath = "Default/Login Data"
	pwdQuery     = `select
	origin_url, username_value, password_value
	from logins`
	outHeader = "uri,service,user,pass,tags\n"
)

var (
	outputFile = "passwords.csv"
	outTags    = "chrome-export"
)

var db *sql.DB

func main() {
	args := os.Args
	if len(args) != 2 {
		p(`Please provide your Chrome configuration folder as argument!`)
		p(`On Windows, this is usually "%LOCALAPPDATA%\Google\Chrome\User Data"`)
		p(`On Linux, this is usually ~/.config/google-chrome`)
		exit(InvalidArgs)
	}
	base := args[1]
	path := path.Join(base, relativePath)
	p("Using " + path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		p("File not found: " + path)
		exit(FileNotFound)
	}

	p("Loading database...")
	db, err := sql.Open("sqlite3", path)
	rows, err := db.Query(pwdQuery)
	if err != nil {
		p(err.Error())
		exit(DatabaseFail)
	}
	cols, err := rows.Columns()
	if err != nil {
		p(err.Error())
		exit(DatabaseFail)
	}
	fmt.Fprint(os.Stderr, "Columns found: ")
	for _, col := range cols {
		fmt.Printf("%s, ", col)
	}
	p("")

	crypt := NewCrypt()

	out, err := os.Create(outputFile)
	if err != nil {
		p("Could not open output file " + outputFile + ": " + err.Error())
		exit(OutputWrite)
	}
	defer out.Close()
	fmt.Fprint(out, outHeader)

	p("Exporting stuff...")
	for rows.Next() {
		var url, username string
		var pwCrypt []byte
		rows.Scan(&url, &username, &pwCrypt)
		password := crypt.decrypt(pwCrypt)
		escapeCSV(&url)
		escapeCSV(&username)
		escapeCSV(&password)
		escapeCSV(&outTags)
		fmt.Fprintf(out, `"%s",,"%s","%s","%s"`+"\n", url, username, password, outTags)
	}
	p("Diddly-done! Your stuff is now in " + outputFile + ` --
- be careful when opening it, make sure nobody is standing behind you!`)
}

func escapeCSV(s *string) {
	*s = strings.Replace(*s, `"`, `""`, -1)
}

func p(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func exit(status Status) {
	os.Exit(int(status))
}
