package main

/*
#cgo LDFLAGS: -lCrypt32
#define WIN32_LEAN_AND_MEAN
#define NOMINMAX
#include <windows.h>
#include <Wincrypt.h>

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
}

void doFree(char* ptr) {
	free(ptr);
}

*/
import "C"

type WindowsCrypt struct {
}

func NewCrypt() crypt {
	return &DarwinCrypt{}
}

func (c *WindowsCrypt) decrypt(input []byte) string {
	var pwLen C.int
	pwDecC := C.decrypt((*C.byte)(&password[0]), C.int(len(password)), &pwLen)
	passwordString := C.GoStringN(pwDecC, pwLen)
	return passwordString
}
