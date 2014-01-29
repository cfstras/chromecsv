package main

/*

https://code.google.com/p/chromium/codesearch#chromium/src/chrome/browser/password_manager/login_database_posix.cc

*/

type LinuxCrypt struct {
}

func NewCrypt() crypt {
	return &LinuxCrypt{}
}

func (c *LinuxCrypt) decrypt(input []byte) string {
	return string(input)
}
