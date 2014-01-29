package main

/*
#cgo LDFLAGS: -lCrypt32
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
	return &WindowsCrypt{}
}

func (c *WindowsCrypt) decrypt(input []byte) string {
	var length C.int
	decruptedC := C.decrypt((*C.byte)(&input[0]), C.int(len(input)), &length)
	decrypted := C.GoStringN(decruptedC, length)
	return decrypted
}
