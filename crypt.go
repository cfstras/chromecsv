package main

type crypt interface {
	decrypt(input []byte) string
}
