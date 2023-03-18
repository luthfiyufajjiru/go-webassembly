package main

import "main/libraries"

func main() {
	c := make(chan struct{})
	defer func() {
		<-c
	}()

	libraries.EncryptFn()

	<-c
}
