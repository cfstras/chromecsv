package main

/*
#cgo LDFLAGS: -lCrypt32
#define WINDOWS_LEAN_AND_MEAN
#include <windows.h>
#include <Wincrypt.h>

char* retval;

char* decrypt(byte* in, int len, int *outLen) {
	DATA_BLOB input, output;
	LPWSTR pDescrOut =  NULL;
	input.cbData = len;
	input.pbData = in;
	CryptUnprotectData(
		&input,
		&pDescrOut,
		NULL,                 // Optional entropy
		NULL,                 // Reserved
		NULL,                 // Here, the optional
							  // prompt structure is not
							  // used.
		0,
		&output);
	*outLen = output.cbData;
	return output.pbData;
	//puts(output.pbData);
}

void doFree(char* ptr) {
	free(ptr);
}

*/
import "C"

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
)

type Status int

const (
	OK Status = iota
	InvalidArgs
	FileNotFound
	DatabaseFail
)

const (
	relativePath = "User Data/Default/Login Data"
	pwdQuery     = `select
	origin_url, username_value, password_value, times_used, password_type
	from logins`
)

var db *sql.DB

func main() {
	args := os.Args
	if len(args) != 2 {
		p(`Please provide your Chrome configuration folder as argument!`)
		p(`On Windows, this is usually %LOCALAPPDATA%\Google\Chrome`)
		p(`On Linux, this is usually ~/.config/Google/Chrome`)
		exit(InvalidArgs)
	}
	base := args[1]
	path := path.Join(base, relativePath)
	p("Using " + path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		p("File not found: " + path)
		exit(FileNotFound)
	}

	fmt.Println("Loading database...")
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
	fmt.Printf("Columns: ")
	for _, col := range cols {
		fmt.Printf("%s, ", col)
	}
	fmt.Println()

	for rows.Next() {
		var url, username, pwType string
		var password []byte
		var timesUsed int
		rows.Scan(&url, &username, &password, &timesUsed, &pwType)
		var pwLen C.int
		pwDecC := C.decrypt((*C.byte)(&password[0]), C.int(len(password)), &pwLen)
		passwordString := C.GoStringN(pwDecC, pwLen)
		//C.doFree(pwDecC)

		fmt.Printf("url: %s, user: %s, pw: %s, pwlen:%d, times used: %d, type: %s\n",
			url, username, passwordString, pwLen, timesUsed, pwType)
		exit(0)
	}

	fmt.Println("Exporting stuff...")
}

func p(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func exit(status Status) {
	os.Exit(int(status))
}
