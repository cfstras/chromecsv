package main

/*
https://code.google.com/p/chromium/codesearch#chromium/src/chrome/browser/password_manager/login_database_mac.cc
https://code.google.com/p/chromium/codesearch#chromium/src/chrome/browser/password_manager/password_store_mac.h
http://dev.chromium.org/developers/design-documents/os-x-password-manager-keychain-integration

//TODO
*/

type DarwinCrypt struct {
}

func NewCrypt() crypt {
	return &DarwinCrypt{}
}

func (c *DarwinCrypt) decrypt(input []byte) string {
	return "Not implemented."
}
